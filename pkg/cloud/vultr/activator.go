package vultr

import (
	_ "github.com/digitalocean/godo"
	_ "golang.org/x/oauth2"

	ol "github.com/jyny/outliner/pkg/outliner"
)

type Activator struct {
	TokenNames []string
	API_TOEKN  string
}

func (a Activator) ListTokenName() []string {
	return a.TokenNames
}

func (a Activator) VerifyToken(token string) bool {
	return false
}

func (a Activator) GenProvider() ol.Provider {
	return *new(Provider)
}
