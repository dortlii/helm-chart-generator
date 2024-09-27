package helm

import (
	"log"
	"os"
	"path"
)

const (
	notesFileName = "Notes.txt"
)

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

// Save implements the function how to save Template to the disk
func (t Template) Save(chartPath string) {
	notesFilePath := path.Join(chartPath, notesFileName)

	if err := os.WriteFile(notesFilePath, []byte(t.Notes.Content), os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
