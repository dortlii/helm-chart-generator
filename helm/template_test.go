package helm

import (
	"reflect"
	"testing"
)

func TestNewTemplate(t *testing.T) {
	tests := []struct {
		name string
		want Template
	}{
		{
			name: "Valid, empty template struct",
			want: Template{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTemplate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemplate_Save(t1 *testing.T) {
	type fields struct {
		Notes Notes
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
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Template{
				Notes: tt.fields.Notes,
			}
			t.Save(tt.args.chartPath)
		})
	}
}
