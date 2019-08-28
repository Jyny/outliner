package digitalocean

import (
	_ "github.com/digitalocean/godo"
	_ "golang.org/x/oauth2"

	ol "github.com/jyny/outliner/pkg/outliner"
)

type DigitalOcean struct {
}

func (do DigitalOcean) TokenKey() []string {
	return []string{
		"",
	}
}

func (do DigitalOcean) Verify([]string) bool {
	return false
}

func (do DigitalOcean) Name() string {
	return "DigitalOcean"
}

func (do DigitalOcean) Region() []string {
	return make([]string, 0)
}

func (do DigitalOcean) ListInstance() []ol.Instance {
	return make([]ol.Instance, 0)
}

func (do DigitalOcean) CreateInstance(ol.InstanceSpec) ol.Instance {
	return ol.Instance{}
}

func (do DigitalOcean) InspectInstance(string) ol.Instance {
	return ol.Instance{}
}

func (do DigitalOcean) DestroyInstance(string) {
}
