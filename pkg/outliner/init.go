package outliner

func init() {
}

// NewCloud generate new cloud
func NewCloud() *Cloud {
	return &Cloud{
		pool: make(map[string]Provider),
	}
}

// NewDeployer generate new deployer
func NewDeployer() *Deployer {
	return &Deployer{}
}
