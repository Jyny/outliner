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

func (p Provider) ListRegion() []ol.Region {
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

func (p Provider) ListSpec() []ol.Spec {
	var ret []ol.Spec
	res, err := p.API.ListTypes(context.Background(), nil)
	if err != nil {
		fmt.Println("List Spec Error", err)
	}
	for _, r := range res {
		if r.Price.Monthly > 50 {
			continue
		}
		s := ol.Spec{
			ID:       r.ID,
			Transfer: fmt.Sprint(r.Transfer),
			Price:    fmt.Sprint(r.Price.Monthly),
		}
		ret = append(ret, s)
	}
	return ret
}

func (p Provider) ListInstance() []ol.Instance {
	return make([]ol.Instance, 0)
}

func (p Provider) CreateInstance(in ol.Instance) ol.Instance {
	res, err := p.API.CreateInstance(
		context.Background(),
		linodego.InstanceCreateOptions{
			Region:         in.Region.ID,
			Type:           in.Spec.ID,
			Tags:           []string{ol.InstanceTag},
			Image:          "linode/ubuntu18.04",
			AuthorizedKeys: []string{},
			RootPass:       "",
		},
	)
	if err != nil {
		fmt.Println("Create Instance Error", err)
	}
	fmt.Println(res)
	return ol.Instance{}
}

func (p Provider) InspectInstance(string) ol.Instance {
	return ol.Instance{}
}

func (p Provider) DestroyInstance(string) {
}
