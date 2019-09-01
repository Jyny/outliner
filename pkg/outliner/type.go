package outliner

// Validater config & verify token
// define valid Provider gen from Activator
type Validater func(Activator) (Provider, error)

// Activator object before generate a Provider
type Activator interface {
	ListTokenName() []string     // list Token names for register
	VerifyToken(string) bool     // verify api key & availability
	GenProvider(string) Provider // Gen a Provider
}

// Provider defin server provider methods
type Provider interface { // new a provider
	Name() string                              // provider's name
	GetToken() string                          // provider's verified Token
	ListSpec() ([]Spec, error)                 // list provider Instance Spec
	ListRegion() ([]Region, error)             // list provider's available regions
	ListInstance() ([]Instance, error)         // list created instance on provider
	CreateInstance(Instance) (Instance, error) // create instance on provider
	WaitInstance(Instance) error               // get info about instance and vpn
	DestroyInstance(string) error              // destroy instance on provider
}

// Agent defin deply Agent methods
type Agent interface {
	GetCredentialPub() string
	Exec(string, string) error
	Watch(string) error
}

// Instance info about server create on server provider
type Instance struct {
	ID       string
	Provider string
	IPv4     string
	Spec     Spec
	Region   Region
	SSHKey   string
	APICert  APICert
}

// Region info about Region
type Region struct {
	ID   string
	Note string
}

// Spec info about server sepc
type Spec struct {
	ID       string
	Transfer string
	Price    string
}

// APICert info about VPN service on instance
type APICert struct {
	APIurl     string
	CertSha256 string
}
