package linode

import (
	ol "github.com/jyny/outliner/pkg/outliner"

	"github.com/linode/linodego"
)

type Linode struct {
	API linodego.Client
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
