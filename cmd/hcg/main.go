package main

import (
	"fmt"
	"log"

	"github.com/dortlii/helm-chart-generator/generator"
)

func main() {
	helmChart, err := generator.NewHelmChart("2", "myChart", "1.0.0", "application")
	if err != nil {
		log.Fatalf("Failed to create helm chart: %s", err)
	}

	emptyHelmChart, err := generator.NewHelmChart("", "", "", "")
	if err != nil {
		log.Fatalf("Failed to create helm chart: %s", err)
	}

	fmt.Printf("This is my helm chart: %s\n", helmChart)
	fmt.Printf("This is my empty helm chart: %s\n", emptyHelmChart)

	// Save to disk
	helm := generator.NewHelmFileService("/tmp/", helmChart)
	helm.Save()
}
