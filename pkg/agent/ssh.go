package agent

import (
	"io/ioutil"
	"log"
	"os/user"
	"path/filepath"
	"strings"

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
			credentialPub: strings.ReplaceAll(string(keypubBytes), "\n", ""),
			credentialPvt: strings.ReplaceAll(string(keypvtBytes), "\n", ""),
		}
	}

	credentialPub, credentialPvt := genNewCredential()
	return SecureShell{
		credentialPub: credentialPub,
		credentialPvt: credentialPvt,
	}
}

func genNewCredential() (string, string) {
	return "", ""
}

func (s SecureShell) GetCredentialPub() string {
	return s.credentialPub
}

func (s SecureShell) Exec(ip string, cmd string) error {
	return nil
}

func (s SecureShell) Watch(ip string) error {
	return nil
}
