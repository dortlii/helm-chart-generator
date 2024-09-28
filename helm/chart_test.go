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
		{
			name: "Invalid defaults",
			args: args{Chart{
				ApiVersion:  "v2",
				Name:        "not-default",
				Type:        "application",
				Description: "A Helm chart for Kubernetes",
				Version:     "0.1.0",
				AppVersion:  "1.16.0",
			}},
			want: &Chart{
				ApiVersion:  "v2",
				Name:        "default",
				Type:        "application",
				Description: "description",
				Version:     "0.0.0",
				AppVersion:  "0.0.0",
			},
		},
		{
			name: "Empty fields",
			args: args{Chart{
				ApiVersion:  "",
				Name:        "",
				Type:        "",
				Description: "",
				Version:     "",
				AppVersion:  "",
			}},
			want: &Chart{
				ApiVersion:  "v2",
				Name:        "default",
				Type:        "application",
				Description: "description",
				Version:     "0.0.0",
				AppVersion:  "0.0.0",
			},
		},
		{
			name: "Incorrect values",
			args: args{Chart{
				ApiVersion:  "v1",
				Name:        "default",
				Type:        "library",
				Description: "A Helm chart for Kubernetes",
				Version:     "0.1.0",
				AppVersion:  "1.16.0",
			}},
			want: &Chart{
				ApiVersion:  "v2",
				Name:        "default",
				Type:        "application",
				Description: "description",
				Version:     "0.0.0",
				AppVersion:  "0.0.0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetDefaults(tt.args.chart); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDefaults() = %v, want %v", got, tt.want)
			}
		})
	}
}
