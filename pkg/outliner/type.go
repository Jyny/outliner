package outliner

type Provider interface {
	Init() bool
	Name() string
	Region() []string
	ListInstance() []Instance
	CreateInstance(InstanceSpec) Instance
	InspectInstance(string) Instance
	DestroyInstance(string)
}

type Instance struct {
	ID string
	InstanceSpec InstanceSpec
	APICert APICert
}

type InstanceSpec struct {
	Spec string
	Region string
	Provider string
}

type APICert struct {
	APIurl string
	CertSha256 string
}