package linode

import (
	"context"
	"net/http"

	"github.com/linode/linodego"
	"golang.org/x/oauth2"
)

type Activator struct {
	TokenNames []string
	API_TOEKN  string
}

func (a Activator) ListTokenName() []string {
	return a.TokenNames
}

func (a Activator) VerifyToken(token string) bool {
	linodeClient := linodego.NewClient(
		&http.Client{
			Transport: &oauth2.Transport{
				Source: oauth2.StaticTokenSource(
					&oauth2.Token{
						AccessToken: token,
					},
				),
			},
		},
	)
	_, err := linodeClient.GetProfile(context.Background())
	linodeClient.SetDebug(true)

	if err != nil {
		return false
	}
	return true
}

func (a Activator) GenClient() *Linode {
	return new(Linode)
}
