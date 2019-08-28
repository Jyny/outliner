package linode

import (
	"context"
	"errors"
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

func (p Provider) ListRegion() ([]ol.Region, error) {
	var ret []ol.Region
	res, err := p.API.ListRegions(context.Background(), nil)
	if err != nil {
		return ret, err
	}
	for _, r := range res {
		t := ol.Region{
			ID:   r.ID,
			Note: r.Country,
		}
		ret = append(ret, t)
	}
	return ret, nil
}

func (p Provider) ListSpec() ([]ol.Spec, error) {
	var ret []ol.Spec
	res, err := p.API.ListTypes(context.Background(), nil)
	if err != nil {
		return ret, err
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
	return ret, nil
}

func (p Provider) ListInstance() ([]ol.Instance, error) {
	var ret []ol.Instance
	res, err := p.API.ListInstances(context.Background(), nil)
	if err != nil {
		return ret, err
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
	return ret, err
}

func (p Provider) CreateInstance(in ol.Instance) (ol.Instance, error) {
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
		return ol.Instance{}, err
	}
	_, err = p.API.WaitForInstanceStatus(
		context.Background(),
		res.ID,
		"running",
		120,
	)
	if err != nil {
		return ol.Instance{}, err
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
	}, nil
}

func (p Provider) InspectInstance(string) (ol.Instance, error) {
	return ol.Instance{}, errors.New("Instance Not Found")
}

func (p Provider) DestroyInstance(id string) error {
	res, err := p.API.ListInstances(context.Background(), nil)
	if err != nil {
		return err
	}

	for _, i := range res {
		if !util.InSliceOfString(i.Tags, ol.InstanceTag) {
			continue
		}
		if i.Label == id {
			err := p.API.DeleteInstance(
				context.Background(),
				i.ID,
			)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("Instance Not Found")
}
