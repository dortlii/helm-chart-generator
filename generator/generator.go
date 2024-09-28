package generator

import "github.com/dortlii/helm-chart-generator/helm"

// NewHelmChart takes parameters to create a helm.Helm object and returns it
func NewHelmChart(apiVersion, name, version, helmType string) (*helm.Helm, error) {
	chart := fillChart(apiVersion, name, version, helmType)
	values := helm.NewValues()
	template := helm.NewTemplate()

	helmChart := helm.Helm{
		Chart:    chart,
		Values:   values,
		Template: template,
	}

	return &helmChart, nil
}

// fillChart creates a new, empty helm.Chart object, fills it with default
// values and overrides them with given values
func fillChart(apiVersion, name, version, helmType string) helm.Chart {
	chart := helm.NewChart()

	chart = *helm.SetDefaults(chart)

	if checkEmptyField(apiVersion) {
		chart.ApiVersion = apiVersion
	}

	if checkEmptyField(name) {
		chart.Name = name
	}

	if checkEmptyField(version) {
		chart.Version = version
	}

	if checkEmptyField(helmType) {
		chart.Type = helmType
	}

	return chart
}

// checkEmptyField checks the given field, if it's empty
// and returns a bool
func checkEmptyField(field string) bool {
	return field != ""
}
