package linode

import (
	"context"
	"errors"
	"fmt"

	"github.com/linode/linodego"

	ol "github.com/jyny/outliner/pkg/outliner"
)

var providerName = "Linode"

type Provider struct {
	verifiedToken string
	API           linodego.Client
}

func (p Provider) Name() string {
	return providerName
}

func (p Provider) GetToken() string {
	return p.verifiedToken
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
		if !ol.InSliceOfString(i.Tags, ol.InstanceTag) {
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
			AuthorizedKeys: []string{in.SSHKey},
			RootPass:       ol.GenRandomPasswd(),
		},
	)
	if err != nil {
		return ol.Instance{}, err
	}

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

func (p Provider) WaitInstance(in ol.Instance) error {
	res, err := p.API.ListInstances(context.Background(), nil)
	if err != nil {
		return err
	}
	for _, i := range res {
		if !ol.InSliceOfString(i.Tags, ol.InstanceTag) {
			continue
		}
		if i.Label == in.ID {
			_, err = p.API.WaitForInstanceStatus(
				context.Background(),
				i.ID,
				"running",
				120,
			)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("Instance Not Found")
}

func (p Provider) DestroyInstance(id string) error {
	res, err := p.API.ListInstances(context.Background(), nil)
	if err != nil {
		return err
	}

	for _, i := range res {
		if !ol.InSliceOfString(i.Tags, ol.InstanceTag) {
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
