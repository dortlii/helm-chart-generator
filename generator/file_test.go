package generator

import (
	"reflect"
	"testing"

	"github.com/dortlii/helm-chart-generator/helm"
)

func TestHelmFiles_Save(t *testing.T) {
	type fields struct {
		Path string
		Helm helm.Helm
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hf := HelmFiles{
				Path: tt.fields.Path,
				Helm: tt.fields.Helm,
			}
			if err := hf.Save(); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewHelmFileService(t *testing.T) {
	type args struct {
		path string
		helm helm.Helm
	}
	tests := []struct {
		name string
		args args
		want FileService
	}{
		{
			name: "Correct HelmFileService",
			args: args{
				path: t.TempDir(),
				helm: helm.Helm{
					Chart: helm.Chart{
						ApiVersion:  "v2",
						Name:        "test",
						Description: "test case",
						Type:        "application",
						Version:     "0.1.0",
						AppVersion:  "0.0.1",
					},
					Values: helm.Values{
						Values: map[string]string{
							"key": "value",
						},
					},
					Template: helm.Template{
						Notes: helm.Notes{
							Content: "This is a test case",
						},
					},
				},
			},
			// TODO: Add want
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHelmFileService(tt.args.path, tt.args.helm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHelmFileService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createFolder(t *testing.T) {
	type args struct {
		folderPath string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createFolder(tt.args.folderPath)
		})
	}
}
