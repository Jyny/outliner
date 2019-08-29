package vultr

import (
	"github.com/vultr/govultr"

	ol "github.com/jyny/outliner/pkg/outliner"
)

var providerName = "Vultr"

type Provider struct {
	API govultr.Client
}

func (p Provider) Name() string {
	return providerName
}

func (p Provider) Region() []string {
	return make([]string, 0)
}

func (p Provider) ListInstance() []ol.Instance {
	return make([]ol.Instance, 0)
}

func (p Provider) CreateInstance(ol.InstanceSpec) ol.Instance {
	return ol.Instance{}
}

func (p Provider) InspectInstance(string) ol.Instance {
	return ol.Instance{}
}

func (p Provider) DestroyInstance(string) {
}
