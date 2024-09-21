package helm

type Template struct {
	Notes Notes
}

type Notes struct {
	Content string
}

// NewTemplate creates an empty template
func NewTemplate() Template {
	return Template{}
}
