package ssh

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/jyny/outliner/pkg/deployer/ssh/consts"
	ol "github.com/jyny/outliner/pkg/outliner"
)

func NewAgent() ol.Agent {
	certok := true
	u, _ := user.Current()

	keypub := filepath.Join(u.HomeDir, consts.SSHKeyPubPath)
	keypvt := filepath.Join(u.HomeDir, consts.SSHKeyPvtPath)

	keypubBytes, err := ioutil.ReadFile(keypub)
	if err != nil {
		certok = false
	}
	keypvtBytes, err := ioutil.ReadFile(keypvt)
	if err != nil {
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

func genNewCredential() (pub string, pvt string) {
	u, _ := user.Current()
	savePrivateFileTo := filepath.Join(u.HomeDir, consts.SSHKeyPvtPath)
	savePublicFileTo := filepath.Join(u.HomeDir, consts.SSHKeyPubPath)
	os.MkdirAll(filepath.Join(u.HomeDir, "/.outliner/"), os.ModePerm)

	bitSize := 4096
	privateKey, err := generatePrivateKey(bitSize)
	if err != nil {
		panic(err)
	}
	publicKeyBytes, err := generatePublicKey(&privateKey.PublicKey)
	if err != nil {
		panic(err)
	}

	privateKeyBytes := encodePrivateKeyToPEM(privateKey)
	err = writeKeyToFile(privateKeyBytes, savePrivateFileTo)
	if err != nil {
		panic(err)
	}
	err = writeKeyToFile([]byte(publicKeyBytes), savePublicFileTo)
	if err != nil {
		panic(err)
	}

	pub = string(publicKeyBytes)
	pvt = string(privateKeyBytes)
	return
}

// generatePrivateKey creates a RSA Private Key of specified byte size
func generatePrivateKey(bitSize int) (*rsa.PrivateKey, error) {
	// Private Key generation
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}

	// Validate Private Key
	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}

	log.Println("[Initializing] Private Key generated")
	return privateKey, nil
}

// encodePrivateKeyToPEM encodes Private Key from RSA to PEM format
func encodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	// Get ASN.1 DER format
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)

	// pem.Block
	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	// Private key in PEM format
	privatePEM := pem.EncodeToMemory(&privBlock)

	return privatePEM
}

// generatePublicKey take a rsa.PublicKey and return bytes suitable for writing to .pub file
// returns in the format "ssh-rsa ..."
func generatePublicKey(privatekey *rsa.PublicKey) ([]byte, error) {
	publicRsaKey, err := ssh.NewPublicKey(privatekey)
	if err != nil {
		return nil, err
	}

	pubKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)

	log.Println("[Initializing] Public key generated")
	return pubKeyBytes, nil
}

// writePemToFile writes keys to a file
func writeKeyToFile(keyBytes []byte, saveFileTo string) error {
	err := ioutil.WriteFile(saveFileTo, keyBytes, 0600)
	if err != nil {
		return err
	}

	log.Printf("[Initializing] Key saved to: %s", saveFileTo)
	return nil
}
