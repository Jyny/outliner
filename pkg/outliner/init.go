package outliner

func init() {
}

func New() *Cloud {
	return &Cloud{make(map[string]Provider)}
}