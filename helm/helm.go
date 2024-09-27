package helm

// Helm chart components
type Helm struct {
	Chart    Chart    `yaml:"chart"`
	Template Template `yaml:"template"`
	Values   Values   `yaml:"values"`
}

// ComponentService exposes file operations
type ComponentService interface {
	Save(path string)
}
