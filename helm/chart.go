package helm

type Chart struct {
	ApiVersion  string
	Name        string
	Description string
	Type        string
	Version     string
	AppVersion  string
}

// New creates an empty helm chart
func NewChart() Chart {
	return Chart{}
}
