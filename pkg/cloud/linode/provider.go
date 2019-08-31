package linode

import (
	"context"
	"fmt"

	"github.com/linode/linodego"

	ol "github.com/jyny/outliner/pkg/outliner"
	"github.com/jyny/outliner/pkg/util"
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
	var ret []ol.Instance
	res, err := p.API.ListInstances(context.Background(), nil)
	if err != nil {
		fmt.Println("List Instances Error", err)
	}

	for _, i := range res {
		if !util.InSliceOfString(i.Tags, ol.InstanceTag) {
			continue
		}
		ret = append(ret, ol.Instance{
			ID:       i.Label,
			Provider: providerName,
			IPv4:     i.IPv4[0].String(),
			Spec: ol.Spec{
				ID: i.Type,
			},
			Region: ol.Region{
				ID: i.Region,
			},
		})
	}

	return ret
}

func (p Provider) CreateInstance(in ol.Instance) ol.Instance {
	res, err := p.API.CreateInstance(
		context.Background(),
		linodego.InstanceCreateOptions{
			Region:         in.Region.ID,
			Type:           in.Spec.ID,
			Tags:           []string{ol.InstanceTag},
			Image:          "linode/ubuntu18.04",
			AuthorizedKeys: []string{util.GetSSHauthorizedKey()},
			RootPass:       util.GenRandomPasswd(),
		},
	)
	if err != nil {
		fmt.Println("Create Instance Error", err)
	}

	// Todo

	return ol.Instance{
		ID:       res.Label,
		Provider: providerName,
		IPv4:     res.IPv4[0].String(),
		Spec: ol.Spec{
			ID: res.Type,
		},
		Region: ol.Region{
			ID: res.Region,
		},
	}
}

func (p Provider) InspectInstance(string) ol.Instance {
	return ol.Instance{}
}

func (p Provider) DestroyInstance(string) {
}
