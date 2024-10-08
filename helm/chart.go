package helm

import (
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

const (
	chartFileName = "Chart.yaml"
)

// Chart item of the helm chart, equals to Chart.yaml
type Chart struct {
	// ApiVersion represents the version of the helm chart schema
	ApiVersion string `yaml:"apiVersion"`
	// Name of the helm chart
	Name string `yaml:"name"`
	// Description describes what the helm chart is for
	Description string `yaml:"description"`
	// Type of the helm chart, application or library
	Type string `yaml:"type"`
	// Version of the helm chart
	Version string `yaml:"version"`
	// AppVersion is the version of the app to be installed
	AppVersion string `yaml:"appVersion"`
}

// NewChart creates an empty helm chart
func NewChart() Chart {
	return Chart{}
}

// SetDefaults adds default values to a Chart struct
func SetDefaults(chart Chart) *Chart {
	chart.ApiVersion = "v2"
	chart.AppVersion = "0.0.0"
	chart.Version = "0.0.0"
	chart.Name = "default"
	chart.Description = "description"
	chart.Type = "application"

	return &chart
}

// Save method for chart component
func (c Chart) Save(chartPath string) {
	yamlData, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	filePath := path.Join(chartPath, chartFileName)

	if err := os.WriteFile(filePath, yamlData, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
