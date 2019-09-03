package agent

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"

	"github.com/jyny/outliner/pkg/agent/consts"
	ol "github.com/jyny/outliner/pkg/outliner"
)

type SecureShell struct {
	credentialPub string
	credentialPvt string
}

func (s SecureShell) GetCredentialPub() string {
	return s.credentialPub
}

func (s SecureShell) sendscript(conn *ssh.Client) error {
	client, err := sftp.NewClient(conn)
	if err != nil {
		return err
	}
	defer client.Close()

	srcFile, err := Script.Open(consts.ScriptName)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := client.Create(filepath.Join("/root/", consts.ScriptName))
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func (s SecureShell) runscript(conn *ssh.Client, cmd string) error {
	session, err := conn.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	if err := session.Run(cmd); err != nil {
		return err
	}

	return nil
}

func (s SecureShell) sshconfig() (*ssh.ClientConfig, error) {
	signer, err := ssh.ParsePrivateKey([]byte(s.credentialPvt))
	if err != nil {
		return &ssh.ClientConfig{}, err
	}

	return &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}, nil
}

func (s SecureShell) sshconn(ip string) (*ssh.Client, error) {
	config, err := s.sshconfig()
	if err != nil {
		return &ssh.Client{}, err
	}
	return ssh.Dial("tcp", ip+":22", config)
}

func (s SecureShell) Deploy(ip string) error {
	conn, err := s.sshconn(ip)
	if err != nil {
		return err
	}

	// sendscript
	err = s.sendscript(conn)
	if err != nil {
		return err
	}
	log.Println("[Deploy script uploaded]")

	// runscript
	cmd := "chmod +x " + consts.ScriptName
	log.Println("[Deploy init]", cmd)
	err = s.runscript(conn, cmd)
	if err != nil {
		return err
	}

	cmd = "nohup bash "
	cmd += filepath.Join("/root/", consts.ScriptName)
	cmd += " &> /tmp/out &"
	log.Println("[Deploy init]", cmd)
	err = s.runscript(conn, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (s SecureShell) Watch(ip string) error {
	conn, err := s.sshconn(ip)
	if err != nil {
		return err
	}

	session, err := conn.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	stdin, err := session.StdinPipe()
	if err != nil {
		return err
	}
	stdout, err := session.StdoutPipe()
	if err != nil {
		return err
	}

	if err := session.Shell(); err != nil {
		return err
	}

	// watch deploy stdout
	cmd := "tail -f /tmp/out"
	_, err = fmt.Fprintf(stdin, "%s\n", cmd)
	if err != nil {
		log.Fatal(err)
	}

	line := make(chan string)
	go func(stdout io.Reader) {
		reader := bufio.NewReader(stdout)
		for {
			linebyte, condition, err := reader.ReadLine()
			if err != nil {
				log.Println(err)
			}
			newline := string(linebyte)
			for condition {
				linebyte, condition, err = reader.ReadLine()
				if err != nil {
					log.Println(err)
				}
				newline += string(linebyte)
			}
			line <- newline
		}
	}(stdout)

	// wait deploy done
	done := make(chan bool)
	go func(ip string) {
		conn, err := s.sshconn(ip)
		if err != nil {
			log.Println(err)
		}

		session, err := conn.NewSession()
		if err != nil {
			log.Println(err)
		}
		defer session.Close()

		cmd := "while kill -0 \"$(cat /tmp/pid)\" &> /dev/null; do sleep 1; done;"
		if err := session.Run(cmd); err != nil {
			log.Println(err)
		}

		done <- true
	}(ip)

	for {
		select {
		case out := <-line:
			log.Println("[Deploy running]", out)
		case <-time.After(1 * time.Minute):
			return errors.New("Watch Deploy Timeout")
		case <-done:
			return nil
		}
	}
}

func (s SecureShell) GetServiceCert(ip string) (ol.APICert, error) {
	conn, err := s.sshconn(ip)
	if err != nil {
		return ol.APICert{}, err
	}

	session, err := conn.NewSession()
	if err != nil {
		return ol.APICert{}, err
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b

	cmd := "cat /opt/outline/access.txt"
	if err := session.Run(cmd); err != nil {
		log.Println(err)
	}

	var ret ol.APICert
	for _, e := range strings.Split(b.String(), "\n") {
		f := strings.Split(e, ":")
		if f[0] == "certSha256" {
			ret.CertSha256 = f[1]
		}
		if f[0] == "apiUrl" {
			ret.APIurl = f[1] + ":" + f[2] + ":" + f[3]
		}
	}
	return ret, nil
}
