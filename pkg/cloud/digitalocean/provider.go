package digitalocean

import (
	ol "github.com/jyny/outliner/pkg/outliner"

	"github.com/digitalocean/godo"
)

type DigitalOcean struct {
	API godo.Client
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
