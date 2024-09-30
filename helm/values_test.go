package helm

import (
	"reflect"
	"testing"
)

func TestNewValues(t *testing.T) {
	tests := []struct {
		name string
		want Values
	}{
		{
			name: "Valid, empty Values struct",
			want: Values{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewValues(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValues_Save(t *testing.T) {
	type fields struct {
		Values map[string]string
	}
	type args struct {
		chartPath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Correctly saved values file",
			fields: fields{
				Values: map[string]string{
					"key": "value",
				},
			},
			args: args{
				chartPath: t.TempDir(),
			},
		},
		{
			name:   "No saved file",
			fields: fields{},
			args: args{
				chartPath: t.TempDir(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Values{
				Values: tt.fields.Values,
			}
			v.Save(tt.args.chartPath)

			// TODO: Check if file is created
		})
	}
}
