package vultr

import (
	ol "github.com/jyny/outliner/pkg/outliner"

	"github.com/vultr/govultr"
)

type Vultr struct {
	API govultr.Client
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
