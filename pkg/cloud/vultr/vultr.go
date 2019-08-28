package vultr

import (
	_ "github.com/vultr/govultr"

	ol "github.com/jyny/outliner/pkg/outliner"
)

type Vultr struct {
}

func (vt Vultr) Init() bool {
	return false
}

func (vt Vultr) Name() string {
	return "Vultr"
}

func (vt Vultr) Region() []string {
	return make([]string, 0)
}

func (vt Vultr) ListInstance() []ol.Instance {
	return make([]ol.Instance, 0)
}

func (vt Vultr) CreateInstance(ol.InstanceSpec) ol.Instance {
	return ol.Instance{}
}

func (vt Vultr) InspectInstance(string) ol.Instance {
	return ol.Instance{}
}

func (vt Vultr) DestroyInstance(string) {
}
