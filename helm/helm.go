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

// ComponentService exposes file operations
type ComponentService interface {
	Save(chartPath string)
}

// Save implements the function how to save Values to the disk
func (h Helm) Save(chartPath string) {
	h.Chart.Save(chartPath)
	h.Values.Save(chartPath)
	h.Template.Save(chartPath)
}
