package outliner

func init() {
}

// New generate new object
func New() *Cloud {
	return &Cloud{make(map[string]Provider)}
}
