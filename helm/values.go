package helm

// Values of the helm chart, equals to a values.yaml file
type Values struct {
	Values []string `yaml:"values"`
}

// NewValues returns a new, empty Values object
func NewValues() Values {
	return Values{}
}
