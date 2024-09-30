package helm

import "testing"

func TestHelm_Save(t *testing.T) {
	type fields struct {
		Chart    Chart
		Values   Values
		Template Template
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
			h := Helm{
				Chart:    tt.fields.Chart,
				Values:   tt.fields.Values,
				Template: tt.fields.Template,
			}
			h.Save(tt.args.chartPath)
		})
	}
}
