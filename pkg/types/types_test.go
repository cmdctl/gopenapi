package types

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	type args struct {
		params Params
	}
	tests := []struct {
		name    string
		args    args
		want    *GeneratedTypes
		want1   []byte
		wantErr bool
	}{
		{
			name:    "should generate types for test/data/v3.0/link-example.yaml",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := Generate(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generate() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Generate() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
