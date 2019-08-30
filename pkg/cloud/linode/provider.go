package linode

import (
	"context"
	"fmt"

	"github.com/linode/linodego"

	ol "github.com/jyny/outliner/pkg/outliner"
)

var providerName = "Linode"

type Provider struct {
	API linodego.Client
}

func (p Provider) Name() string {
	return providerName
}

func (p Provider) Region() []ol.Region {
	var ret []ol.Region
	res, err := p.API.ListRegions(context.Background(), nil)
	if err != nil {
		fmt.Println("Finding Region Error", err)
	}
	for _, r := range res {
		t := ol.Region{
			ID:   r.ID,
			Note: r.Country,
		}
		ret = append(ret, t)
	}
	return ret
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
