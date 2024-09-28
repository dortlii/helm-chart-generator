package helm

import (
	"reflect"
	"testing"
)

func TestChart_Save(t *testing.T) {
	type fields struct {
		ApiVersion  string
		Name        string
		Description string
		Type        string
		Version     string
		AppVersion  string
	}
	type args struct {
		chartPath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Chart{
				ApiVersion:  tt.fields.ApiVersion,
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				Type:        tt.fields.Type,
				Version:     tt.fields.Version,
				AppVersion:  tt.fields.AppVersion,
			}
			c.Save(tt.args.chartPath)
		})
	}
}

func TestNewChart(t *testing.T) {
	tests := []struct {
		name string
		want Chart
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChart(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetDefaults(t *testing.T) {
	type args struct {
		chart Chart
	}
	tests := []struct {
		name string
		args args
		want *Chart
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetDefaults(tt.args.chart); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDefaults() = %v, want %v", got, tt.want)
			}
		})
	}
}
