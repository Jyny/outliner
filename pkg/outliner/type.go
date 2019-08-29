package outliner

// Validater define valid Provider to gen from Activator
type Validater func(Activator) (Provider, error)

// Activator object before generate a Provider
type Activator interface {
	ListTokenName() []string // list Token names for register
	VerifyToken(string) bool // verify api key & availability
	GenProvider() Provider   // Gen a Provider
}

// Provider defin server provider methods
type Provider interface { // new a provider
	Name() string                         // provider's name
	Region() []string                     // list provider's available regions
	ListInstance() []Instance             // list created instance on provider
	CreateInstance(InstanceSpec) Instance // create instance on provider
	InspectInstance(string) Instance      // get info about instance and vpn
	DestroyInstance(string)               // destroy instance on provider
}

// Instance info about server create on server provider
type Instance struct {
	ID           string
	InstanceSpec InstanceSpec
	APICert      APICert
}

// InstanceSpec info about server sepc
type InstanceSpec struct {
	Spec     string
	Region   string
	Provider string
}

// APICert info about VPN service on instance
type APICert struct {
	APIurl     string
	CertSha256 string
}
