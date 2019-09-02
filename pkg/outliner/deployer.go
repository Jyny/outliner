package outliner

import ()

// Deployer core object for outliner
type Deployer struct {
	agent Agent
}

// Init deployer cert init
func (d *Deployer) Init(agent Agent) {
	d.agent = agent
}

// GetCredentialPub Get CredentialPub
func (d Deployer) GetCredentialPub() string {
	return d.agent.GetCredentialPub()
}

// DeployService deploy service
func (d Deployer) DeployService(ip string) error {
	return d.agent.Deploy(ip)
}

// WaitService wait service
func (d Deployer) WaitService(ip string) error {
	return d.agent.Watch(ip)
}

// GetServiceCert get service cert
func (d Deployer) GetServiceCert(ip string) (APICert, error) {
	return d.agent.GetServiceCert(ip)
}
