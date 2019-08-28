package outliner

type Cloud struct {
	pool map[string]Provider
}

func (c *Cloud)AddProvider(prvders ...Provider) {
	for _, prvder := range prvders {
		if prvder.Init() {
			c.pool[prvder.Name()] = prvder
		}
	}
}

func (c *Cloud)LookupRegion() map[string][]string{
	mapPrvderRegion := make( map[string][]string)
	for _, prvder := range c.pool {
		var regs []string
		for _, reg := range prvder.Region() {
			regs = append(regs, reg)
		}
		mapPrvderRegion[prvder.Name()] = regs
	}
	return mapPrvderRegion
}

func (c *Cloud)ListInstance() []Instance {
	var ret []Instance
	for _, prvder := range c.pool {
		for _, inst := range prvder.ListInstance() {
			ret = append(ret, inst)
		}
	}
	return ret
}

func (c *Cloud)CreateInstance(spec InstanceSpec) Instance {
	return c.pool[spec.Provider].CreateInstance(spec)
}

func (c *Cloud)InspectInstance(ID string) Instance {
	for _, prvder := range c.pool {
		for _, inst := range prvder.ListInstance() {
			if inst.ID == ID {
				return c.pool[inst.InstanceSpec.Provider].InspectInstance(ID)
			}
		}
	}
	return Instance{}
}

func (c *Cloud)DestroyInstance(ID string) {
	for _, prvder := range c.pool {
		for _, inst := range prvder.ListInstance() {
			if inst.ID == ID {
				c.pool[inst.InstanceSpec.Provider].DestroyInstance(ID)
			}
		}
	}
}