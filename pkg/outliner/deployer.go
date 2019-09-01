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
	return d.agent.Exec(ip, "whoami")
}

// WaitService wait service
func (d Deployer) WaitService(string) error {
	return nil
}

// GetServiceCert get service cert
func (d Deployer) GetServiceCert(string) (APICert, error) {
	return APICert{}, nil
}
