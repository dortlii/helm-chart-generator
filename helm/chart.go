package helm

type Chart struct {
	// ApiVersion represents the version of the helm chart schema
	ApiVersion string
	// Name of the helm chart
	Name string
	// Description describes what the helm chart is for
	Description string
	// Type of the helm chart, application or library
	Type string
	// Version of the helm chart
	Version string
	// AppVersion is the version of the app to be installed
	AppVersion string
}

// NewChart creates an empty helm chart
func NewChart() Chart {
	return Chart{}
}

// SetDefaults adds default values to a Chart struct
func SetDefaults(chart Chart) *Chart {
	chart.ApiVersion = "2"
	chart.AppVersion = "0.0.0"
	chart.Version = "0.0.0"
	chart.Name = "default"
	chart.Description = "description"
	chart.Type = "application"

	return &chart
}
