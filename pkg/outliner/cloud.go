package outliner

// Cloud core object for outliner
type Cloud struct {
	pool map[string]Provider
}

// RegisterProvider Register a cloud Provider whith Validater function
func (c *Cloud) RegisterProvider(validater Validater, actvrs ...Activator) {
	for _, actvr := range actvrs {
		prvdr, err := validater(actvr)
		if err != nil {
			continue
		}
		c.pool[prvdr.Name()] = prvdr
	}
}

// LookupSpec show avalible Specs on Providers
func (c *Cloud) LookupSpec() map[string][]InstanceSpec {
	ret := make(map[string][]InstanceSpec)
	for _, prvder := range c.pool {
		var specs []InstanceSpec
		for _, spec := range prvder.ListSpec() {
			specs = append(specs, spec)
		}
		ret[prvder.Name()] = specs
	}
	return ret
}

// LookupRegion show avalible Regions on Providers
func (c *Cloud) LookupRegion() map[string][]Region {
	ret := make(map[string][]Region)
	for _, prvder := range c.pool {
		var regs []Region
		for _, reg := range prvder.ListRegion() {
			regs = append(regs, reg)
		}
		ret[prvder.Name()] = regs
	}
	return ret
}

// ListInstance list all instances create by outliner
func (c *Cloud) ListInstance() []Instance {
	var ret []Instance
	for _, prvder := range c.pool {
		for _, inst := range prvder.ListInstance() {
			ret = append(ret, inst)
		}
	}
	return ret
}

// CreateInstance create a instance on server Provider
func (c *Cloud) CreateInstance(in Instance) Instance {
	return c.pool[in.Provider].CreateInstance(in)
}

// InspectInstance show the instance and VPN service info
func (c *Cloud) InspectInstance(ID string) Instance {
	for _, prvder := range c.pool {
		for _, inst := range prvder.ListInstance() {
			if inst.ID == ID {
				return c.pool[inst.Provider].InspectInstance(ID)
			}
		}
	}
	return Instance{}
}

// DestroyInstance destroy a instanceon server Provider
func (c *Cloud) DestroyInstance(ID string) {
	for _, prvder := range c.pool {
		for _, inst := range prvder.ListInstance() {
			if inst.ID == ID {
				c.pool[inst.Provider].DestroyInstance(ID)
			}
		}
	}
}
