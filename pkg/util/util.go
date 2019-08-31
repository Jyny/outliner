package util

import (
	"io/ioutil"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/sethvargo/go-password/password"
)

func InSliceOfString(s []string, x string) bool {
	for _, e := range s {
		if x == e {
			return true
		}
	}
	return false
}

func GetSSHauthorizedKey() string {
	// Read authorized ssh key
	u, _ := user.Current()
	key := filepath.Join(u.HomeDir, "/.ssh/id_rsa.pub")
	authorizedKeysBytes, err := ioutil.ReadFile(key)
	if err != nil {
		panic(err)
	}
	authorizedKey := strings.ReplaceAll(string(authorizedKeysBytes), "\n", "")

	return authorizedKey
}

func GenRandomPasswd() string {
	// Generate a password that is 64 characters long with 10 digits, 10 symbols,
	// allowing upper and lower case letters, disallowing repeat characters.
	rpwd, err := password.Generate(64, 10, 10, false, false)
	if err != nil {
		panic(err)
	}
	return rpwd
}
