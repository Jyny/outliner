package util

import (
	"github.com/sethvargo/go-password/password"
	"time"
)

func InSliceOfString(s []string, x string) bool {
	for _, e := range s {
		if x == e {
			return true
		}
	}
	return false
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

func Waitforawhile() {
	time.Sleep(15 * time.Second)
}
