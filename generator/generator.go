package generator

import "github.com/dortlii/helm-chart-generator/helm"

func NewHelmChart(apiVersion, name, version, helmType string) (helm.Helm, error) {
	chart := fillChart(apiVersion, name, version, helmType)
	values := helm.NewValues()
	template := helm.NewTemplate()

	helm := helm.Helm{
		Chart:    chart,
		Values:   values,
		Template: template,
	}

	return helm, nil
}

func fillChart(apiVersion, name, version, helmType string) helm.Chart {
	chart := helm.NewChart()
	chart.ApiVersion = apiVersion
	chart.Name = name
	chart.Version = version
	chart.Type = helmType

	return chart
}
