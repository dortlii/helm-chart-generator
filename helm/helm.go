package helm

// Helm chart components
type Helm struct {
	// Chart of the helm chart
	Chart Chart
	// Values of the helm chart
	Values Values
	// Template of the helm chart
	Template Template
}
