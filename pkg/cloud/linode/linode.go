package linode

import (
	_ "golang.org/x/oauth2"
	_ "github.com/linode/linodego"
	
	ol "github.com/jyny/outliner/pkg/outliner"
)

type Linode struct {
}

func (li Linode)Init() bool {
	return true
}

func (li Linode)Name() string {
	return "Linode"
}

func (li Linode)Region() []string {
	return make([]string, 0)
}

func (li Linode)ListInstance() []ol.Instance {
	return make([]ol.Instance, 0)
}

func (li Linode)CreateInstance(ol.InstanceSpec) ol.Instance {
	return ol.Instance{}
}

func (li Linode)InspectInstance(string) ol.Instance {
	return ol.Instance{}
}

func (li Linode)DestroyInstance(string) {
}