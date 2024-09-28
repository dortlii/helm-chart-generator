package main

import (
	"fmt"
	"log"

	"github.com/dortlii/helm-chart-generator/generator"
)

func main() {
	helmChart, err := generator.NewHelmChart("v2", "myChart", "1.0.0", "application")
	if err != nil {
		log.Fatalf("Failed to create helm chart: %s", err)
	}

	// Add some text to the Template Notes for testing
	helmChart.Template.Notes.Content = "I'm a note!"

	// Add some values to Values for testing
	helmChart.Values.Values = map[string]string{
		"key":        "value",
		"anotherKey": "anotherValue",
	}

	emptyHelmChart, err := generator.NewHelmChart("", "", "", "")
	if err != nil {
		log.Fatalf("Failed to create helm chart: %s", err)
	}

	fmt.Printf("This is my helm chart: %s\n", helmChart)
	fmt.Printf("This is my empty helm chart: %s\n", emptyHelmChart)

	// Save to disk
	hfs := generator.NewHelmFileService("/tmp/", helmChart)
	hfs.Save()
}
