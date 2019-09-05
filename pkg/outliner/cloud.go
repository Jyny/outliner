package outliner

import (
	"errors"
)

// Cloud core object for outliner
type Cloud struct {
	pool map[string]Provider
}

// RegisterProvider Register a cloud Provider whith Validater function
func (c *Cloud) RegisterProvider(validater Validater, actvrs ...Activator) error {
	if len(actvrs) == 0 {
		return errors.New("No avalible Provider")
	}
	for _, actvr := range actvrs {
		prvdr, err := validater(actvr)
		if err != nil {
			continue
		}
		c.pool[prvdr.Name()] = prvdr
	}
	return nil
}

// ListSpec show avalible Specs on Providers
func (c *Cloud) ListSpec() (map[string][]Spec, error) {
	ret := make(map[string][]Spec)
	for _, prvder := range c.pool {
		specs, err := prvder.ListSpec()
		if err != nil {
			return ret, err
		}
		ret[prvder.Name()] = specs
	}
	return ret, nil
}

// ListRegion show avalible Regions on Providers
func (c *Cloud) ListRegion() (map[string][]Region, error) {
	ret := make(map[string][]Region)
	for _, prvder := range c.pool {
		regs, err := prvder.ListRegion()
		if err != nil {
			return ret, err
		}
		ret[prvder.Name()] = regs
	}
	return ret, nil
}

// ListProvider show avalible Providers
func (c *Cloud) ListProvider() ([][]string, error) {
	var ret [][]string
	if len(c.pool) == 0 {
		return ret, errors.New("No avalible Provider")
	}
	for _, prvder := range c.pool {
		ret = append(ret, []string{
			prvder.Name(),
			prvder.GetToken(),
		})
	}
	return ret, nil
}

// ListInstance list all instances create by outliner
func (c *Cloud) ListInstance() ([]Instance, error) {
	var ret []Instance
	for _, prvder := range c.pool {
		insts, err := prvder.ListInstance()
		if err != nil {
			return ret, err
		}
		for _, inst := range insts {
			ret = append(ret, inst)
		}
	}
	return ret, nil
}

// CreateInstance create a instance on server Provider
func (c *Cloud) CreateInstance(in Instance) (Instance, error) {
	return c.pool[in.Provider].CreateInstance(in)
}

// WaitInstance wait instance to boot
func (c Cloud) WaitInstance(in Instance) error {
	return c.pool[in.Provider].WaitInstance(in)
}

// InspectInstanceByIP Inspect Instance
func (c *Cloud) InspectInstanceByIP(IP string) (Instance, error) {
	for _, prvder := range c.pool {
		insts, err := prvder.ListInstance()
		if err != nil {
			return Instance{}, err
		}
		for _, inst := range insts {
			if inst.IPv4 == IP {
				return inst, nil
			}
		}
	}
	return Instance{}, errors.New("Instance Not Found")
}

// InspectInstanceByID Inspect Instance
func (c *Cloud) InspectInstanceByID(ID string) (Instance, error) {
	for _, prvder := range c.pool {
		insts, err := prvder.ListInstance()
		if err != nil {
			return Instance{}, err
		}
		for _, inst := range insts {
			if inst.ID == ID {
				return inst, nil
			}
		}
	}
	return Instance{}, errors.New("Instance Not Found")
}

// DestroyInstance destroy a instanceon server Provider
func (c *Cloud) DestroyInstance(ID string) error {
	for _, prvder := range c.pool {
		insts, err := prvder.ListInstance()
		if err != nil {
			return err
		}
		for _, inst := range insts {
			if inst.ID == ID {
				return c.pool[inst.Provider].DestroyInstance(ID)
			}
		}
	}
	return errors.New("Instance Not Found")
}
