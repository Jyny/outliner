package linode

import (
	"context"
	"net/http"

	"github.com/linode/linodego"
	"golang.org/x/oauth2"

	ol "github.com/jyny/outliner/pkg/outliner"
)

var tokenNames = []string{
	"LINODE_TOKEN",
	"LINODE_CLI_TOKEN",
	"LINODE_API_TOKEN",
}

type Activator struct {
}

func (a Activator) ListTokenName() []string {
	return tokenNames
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

	if err != nil {
		return false
	}
	return true
}

func (a Activator) GenProvider(verifiedToken string) ol.Provider {
	return Provider{
		verifiedToken: verifiedToken,
		API: linodego.NewClient(
			&http.Client{
				Transport: &oauth2.Transport{
					Source: oauth2.StaticTokenSource(
						&oauth2.Token{
							AccessToken: verifiedToken,
						},
					),
				},
			},
		),
	}
}
