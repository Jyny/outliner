package vultr

import (
	_ "github.com/digitalocean/godo"
	_ "golang.org/x/oauth2"

	ol "github.com/jyny/outliner/pkg/outliner"
)

var tokenNames = []string{
	"TOKEN",
}

type Activator struct {
}

func (a Activator) ListTokenName() []string {
	return tokenNames
}

func (a Activator) VerifyToken(token string) bool {
	return false
}

func (a Activator) GenProvider() ol.Provider {
	return *new(Provider)
}
