package core

import (
	"reflect"
	"testing"
)

func TestReadConfig(t *testing.T) {
	type args struct {
		config string
	}
	tests := []struct {
		name    string
		args    args
		want    *GeneratorConfig
		wantErr bool
	}{
		{
			name: "TestReadConfigFromYaml",
			args: args{
				config: "../../gorm-gen.yml",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadConfig(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
