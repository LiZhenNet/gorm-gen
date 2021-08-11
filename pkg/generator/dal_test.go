package generator

import (
	"github.com/lizhennet/gorm-gen/pkg/core"
	"testing"
)

func TestGenDalFile(t *testing.T) {
	type args struct {
		ctx core.Context
	}
	ctx, _ := NewGenCtx("../../gorm-gen.yml", "project")
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "TestGenDalFile",
			args:    args{ctx: ctx},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenDalFile(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("GenDalFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
