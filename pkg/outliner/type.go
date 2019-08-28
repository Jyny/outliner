package outliner

type Provider interface {
	TokenKey() []string                   // list Token key for register
	Verify([]string) bool                 // verify api key & availability & init
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
