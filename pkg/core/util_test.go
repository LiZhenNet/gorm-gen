package core

import "testing"

func TestGetPackageName(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestGetPackageName",
			args: args{
				path: "./output/model",
			},
			want: "model",
		},
		{
			name: "TestGetPackageName",
			args: args{
				path: "/demo/model",
			},
			want: "model",
		},
		{
			name: "TestGetPackageName",
			args: args{
				path: "",
			},
			want: ".",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetPackageName(tt.args.path)
			if got != tt.want {
				t.Errorf("GetPackageName() = %v, want %v", got, tt.want)
			}
		})
	}
}
