package database

import (
	"context"
	"reflect"
	"testing"

	"github.com/lizhennet/gorm-gen/pkg/core"
)

var connectionConfig = core.ConnectionConfig{
	Host:     "127.0.0.1",
	Port:     "33066",
	User:     "root",
	Passport: "123456",
	Database: "demo",
}

func TestGetDbConnection(t *testing.T) {
	type args struct {
		ctx              context.Context
		connectionConfig core.ConnectionConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestGetDbConnectionSuccess",
			args: args{
				ctx:              context.Background(),
				connectionConfig: connectionConfig,
			},
			wantErr: false,
		},
		{
			name: "TestGetDbConnectionError",
			args: args{
				ctx: context.Background(),
				connectionConfig: core.ConnectionConfig{
					Host:     "127.0.0.1",
					Port:     "33066",
					User:     "root",
					Passport: "1234567",
					Database: "demo",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetDbConnection(tt.args.connectionConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDbConnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetTableSchema(t *testing.T) {
	type args struct {
		ctx              context.Context
		connectionConfig core.ConnectionConfig
		tableName        string
	}
	tests := []struct {
		name    string
		args    args
		want    []core.ColumnMeta
		wantErr bool
	}{
		{
			name: "TestGetTableSchema",
			args: args{
				ctx:              context.Background(),
				connectionConfig: connectionConfig,
				tableName:        "project",
			},
			want: []core.ColumnMeta{{
				ColumnName:        "id",
				ColumnDescription: "id",
				ColumnType:        "bigint",
				IsNullAble:        false,
				IsPrimary:         true,
			},
				{
					ColumnName:        "name",
					ColumnDescription: "project name",
					ColumnType:        "varchar",
					IsNullAble:        false,
					IsPrimary:         false,
				},
				{
					ColumnName:        "create_time",
					ColumnDescription: "project create time",
					ColumnType:        "timestamp",
					IsNullAble:        false,
					IsPrimary:         false,
				}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTableSchema(tt.args.connectionConfig, tt.args.tableName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTableSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTableSchema() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTableSchema1(t *testing.T) {
	type args struct {
		connectionConfig core.ConnectionConfig
		tableName        string
	}
	tests := []struct {
		name    string
		args    args
		want    []core.ColumnMeta
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTableSchema(tt.args.connectionConfig, tt.args.tableName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTableSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTableSchema() got = %v, want %v", got, tt.want)
			}
		})
	}
}
