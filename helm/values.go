package helm

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

const (
	valuesFileName = "values.yaml"
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
func (v Values) Save(chartPath string) {
	jsonData, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}

	filePath := path.Join(chartPath, valuesFileName)

	if err := os.WriteFile(filePath, jsonData, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
