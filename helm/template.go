package helm

import (
	"encoding/json"
	"log"
	"os"
)

// Template elements of the helm chart
type Template struct {
	// Notes content for the Template
	Notes Notes `yaml:"notes"`
}

// Notes for the helm chart functionality
type Notes struct {
	// Content describes the text which gets written to Notes
	Content string `yaml:"content"`
}

// NewTemplate creates an empty template
func NewTemplate() Template {
	return Template{}
}

// Save implements the function how to save Template to the disk
func (t Template) Save(path string) {
	jsonData, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(path, jsonData, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
