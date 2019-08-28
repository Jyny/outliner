package linode

import (
	"context"
	"net/http"

	"github.com/linode/linodego"
	"golang.org/x/oauth2"

	ol "github.com/jyny/outliner/pkg/outliner"
)

type Linode struct {
	API_KEY string
}

func (li Linode) TokenKey() []string {
	return []string{
		"LINODE_TOKEN",
		"LINODE_API_KEY",
	}
}

func (li Linode) Verify(apikeys []string) bool {
	linodeClient := linodego.NewClient(
		&http.Client{
			Transport: &oauth2.Transport{
				Source: oauth2.StaticTokenSource(
					&oauth2.Token{
						AccessToken: "",
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

func (li Linode) Name() string {
	return "Linode"
}

func (li Linode) Region() []string {
	return make([]string, 0)
}

func (li Linode) ListInstance() []ol.Instance {
	return make([]ol.Instance, 0)
}

func (li Linode) CreateInstance(ol.InstanceSpec) ol.Instance {
	return ol.Instance{}
}

func (li Linode) InspectInstance(string) ol.Instance {
	return ol.Instance{}
}

func (li Linode) DestroyInstance(string) {
}
