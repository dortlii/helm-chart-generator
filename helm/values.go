package helm

import (
	"encoding/json"
	"log"
	"os"
)

// Values of the helm chart, equals to a values.yaml file
type Values struct {
	Values []string `yaml:"values"`
}

// NewValues returns a new, empty Values object
func NewValues() Values {
	return Values{}
}

// Save implements the function how to save Values to the disk
func (v Values) Save(path string) {
	jsonData, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(path, jsonData, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
