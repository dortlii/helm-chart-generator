package generator

import (
	"reflect"
	"testing"

	"github.com/dortlii/helm-chart-generator/helm"
)

func TestNewHelmChart(t *testing.T) {
	type args struct {
		apiVersion string
		name       string
		version    string
		helmType   string
	}
	tests := []struct {
		name    string
		args    args
		want    *helm.Helm
		wantErr bool
	}{
		{
			name: "Valid Helm struct with default values",
			args: args{
				apiVersion: "v2",
				name:       "default",
				version:    "0.0.0",
				helmType:   "application",
			},
			want: &helm.Helm{
				Chart: helm.Chart{
					ApiVersion:  "v2",
					Name:        "default",
					Description: "description",
					Type:        "application",
					Version:     "0.0.0",
					AppVersion:  "0.0.0",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewHelmChart(tt.args.apiVersion, tt.args.name, tt.args.version, tt.args.helmType)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHelmChart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHelmChart() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkEmptyField(t *testing.T) {
	type args struct {
		field string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkEmptyField(tt.args.field); got != tt.want {
				t.Errorf("checkEmptyField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fillChart(t *testing.T) {
	type args struct {
		apiVersion string
		name       string
		version    string
		helmType   string
	}
	tests := []struct {
		name string
		args args
		want helm.Chart
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fillChart(tt.args.apiVersion, tt.args.name, tt.args.version, tt.args.helmType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fillChart() = %v, want %v", got, tt.want)
			}
		})
	}
}
