package outliner

type Activator interface {
	ListTokenName() []string // list Token names for register
	VerifyToken(string) bool // verify api key & availability
	GenProvider() *Provider  // Gen a Provider
}

type Provider interface { // new a provider
	Name() string                         // provider's name
	Region() []string                     // list provider's available regions
	ListInstance() []Instance             // list created instance on provider
	CreateInstance(InstanceSpec) Instance // create instance on provider
	InspectInstance(string) Instance      // get info about instance and vpn
	DestroyInstance(string)               // destroy instance on provider
}

type Instance struct {
	ID           string
	InstanceSpec InstanceSpec
	APICert      APICert
}

type InstanceSpec struct {
	Spec     string
	Region   string
	Provider string
}

type APICert struct {
	APIurl     string
	CertSha256 string
}
