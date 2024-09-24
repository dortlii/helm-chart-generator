package helm

// Helm chart components
type Helm struct {
	// Chart of the helm chart
	Chart Chart `yaml:"chart"`
	// Values of the helm chart
	Values Values `yaml:"values"`
	// Template of the helm chart
	Template Template `yaml:"template"`
}
