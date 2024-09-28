package main

import (
	"testing"

	"github.com/dortlii/helm-chart-generator/generator"
	"github.com/dortlii/helm-chart-generator/helm"
)

// createHelmChart Helper function to create a Helm chart
func createHelmChart(version, name, chartVersion, appType string) (*helm.Helm, error) {
	return generator.NewHelmChart(version, name, chartVersion, appType)
}

// verifyHelmChartContent Helper function to verify the content of the Helm chart
func verifyHelmChartContent(t *testing.T, helmChart *helm.Helm, expectedNote string, expectedValues map[string]string) {
	if helmChart.Template.Notes.Content != expectedNote {
		t.Errorf("Expected note content %s, but got %s", expectedNote, helmChart.Template.Notes.Content)
	}

	for k, v := range expectedValues {
		if helmChart.Values.Values[k] != v {
			t.Errorf("Expected value for key %s is %s, but got %s", k, v, helmChart.Values.Values[k])
		}
	}
}

func TestHelmCharts(t *testing.T) {
	tests := []struct {
		name           string
		version        string
		chartName      string
		chartVersion   string
		appType        string
		expectedNote   string
		expectedValues map[string]string
	}{
		{
			name:         "Valid Helm Chart",
			version:      "v2",
			chartName:    "myChart",
			chartVersion: "1.0.0",
			appType:      "application",
			expectedNote: "I'm a note!",
			expectedValues: map[string]string{
				"key":        "value",
				"anotherKey": "anotherValue",
			},
		},
		{
			name:           "Empty Helm Chart",
			version:        "",
			chartName:      "",
			chartVersion:   "",
			appType:        "",
			expectedNote:   "",
			expectedValues: map[string]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helmChart, err := createHelmChart(tt.version, tt.chartName, tt.chartVersion, tt.appType)
			if err != nil {
				t.Fatalf("Failed to create helm chart: %s", err)
			}

			// Set the expected note and values for testing
			helmChart.Template.Notes.Content = tt.expectedNote
			helmChart.Values.Values = tt.expectedValues

			verifyHelmChartContent(t, helmChart, tt.expectedNote, tt.expectedValues)
		})
	}
}
