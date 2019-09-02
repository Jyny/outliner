package agent

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"

	ol "github.com/jyny/outliner/pkg/outliner"
)

type SecureShell struct {
	credentialPub string
	credentialPvt string
}

func New() ol.Agent {
	certok := true
	u, _ := user.Current()

	keypub := filepath.Join(u.HomeDir, "/.ssh/id_rsa.pub")
	keypvt := filepath.Join(u.HomeDir, "/.ssh/id_rsa")

	keypubBytes, err := ioutil.ReadFile(keypub)
	if err != nil {
		log.Println(err)
		certok = false
	}
	keypvtBytes, err := ioutil.ReadFile(keypvt)
	if err != nil {
		log.Println(err)
		certok = false
	}

	if certok {
		return SecureShell{
			credentialPub: strings.TrimRight(string(keypubBytes), "\n"),
			credentialPvt: strings.TrimRight(string(keypvtBytes), "\n"),
		}
	}

	credentialPub, credentialPvt := genNewCredential()
	return SecureShell{
		credentialPub: credentialPub,
		credentialPvt: credentialPvt,
	}
}

func genNewCredential() (string, string) {
	// Todo
	return "", ""
}

func (s SecureShell) GetCredentialPub() string {
	return s.credentialPub
}

func sendscript(conn *ssh.Client) error {
	client, err := sftp.NewClient(conn)
	if err != nil {
		return err
	}
	defer client.Close()

	srcFile, err := Script.Open(ScriptName)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := client.Create(filepath.Join("/root/", ScriptName))
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

func runscript(conn *ssh.Client, cmd string) error {
	session, err := conn.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b

	if err := session.Run(cmd); err != nil {
		return err
	}

	return nil
}

func (s SecureShell) Deploy(ip string) error {
	signer, err := ssh.ParsePrivateKey([]byte(s.credentialPvt))
	if err != nil {
		return err
	}

	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", ip+":22", config)
	if err != nil {
		return err
	}

	// sendscript
	err = sendscript(conn)
	if err != nil {
		return err
	}
	log.Println("[Deploy script uploaded]")

	// runscript
	cmd := "chmod +x " + ScriptName
	log.Println("[Running]", cmd)
	err = runscript(conn, cmd)
	if err != nil {
		return err
	}

	cmd = "nohup bash " + filepath.Join("/root/", ScriptName) + " &> /tmp/out &"
	log.Println("[Running]", cmd)
	err = runscript(conn, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (s SecureShell) Watch(ip string) error {
	signer, err := ssh.ParsePrivateKey([]byte(s.credentialPvt))
	if err != nil {
		return err
	}

	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", ip+":22", config)
	if err != nil {
		return err
	}

	session, err := conn.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("whoami"); err != nil {
		return err
	}

	log.Println(b.String())

	return nil
}
