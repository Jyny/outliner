package digitalocean

import (
	_ "context"

	_ "github.com/digitalocean/godo"
	_ "golang.org/x/oauth2"
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

func (a Activator) GenClient() *DigitalOcean {
	return new(DigitalOcean)
}
