package helm

// Template elements of the helm chart
type Template struct {
	// Notes content for the Template
	Notes Notes
}

// Notes for the helm chart functionality
type Notes struct {
	// Content describes the text which gets written to Notes
	Content string
}

// NewTemplate creates an empty template
func NewTemplate() Template {
	return Template{}
}
