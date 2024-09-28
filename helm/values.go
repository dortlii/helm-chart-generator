package helm

import (
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

const (
	valuesFileName = "values.yaml"
)

// Values of the helm chart, equals to a values.yaml file
type Values struct {
	Values map[string]string `yaml:"values"`
}

// NewValues returns a new, empty Values object
func NewValues() Values {
	return Values{}
}

// Save implements the function how to save Values to the disk
func (v Values) Save(chartPath string) {
	yamlData, err := yaml.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}

	filePath := path.Join(chartPath, valuesFileName)

	if err := os.WriteFile(filePath, yamlData, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
