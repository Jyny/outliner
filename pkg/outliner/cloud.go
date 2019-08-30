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

// LookupRegion show avalible Regions on Providers
func (c *Cloud) LookupRegion() map[string][]string {
	mapPrvderRegion := make(map[string][]string)
	for _, prvder := range c.pool {
		var regs []string
		for _, reg := range prvder.Region() {
			regs = append(regs, reg)
		}
		mapPrvderRegion[prvder.Name()] = regs
	}
	return mapPrvderRegion
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
func (c *Cloud) CreateInstance(spec InstanceSpec) Instance {
	return c.pool[spec.Provider].CreateInstance(spec)
}

// InspectInstance show the instance and VPN service info
func (c *Cloud) InspectInstance(ID string) Instance {
	for _, prvder := range c.pool {
		for _, inst := range prvder.ListInstance() {
			if inst.ID == ID {
				return c.pool[inst.InstanceSpec.Provider].InspectInstance(ID)
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
				c.pool[inst.InstanceSpec.Provider].DestroyInstance(ID)
			}
		}
	}
}
