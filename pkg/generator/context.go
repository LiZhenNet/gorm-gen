package generator

import (
	"errors"

	"github.com/lizhennet/gorm-gen/pkg/core"
	"github.com/lizhennet/gorm-gen/pkg/database"
)

func NewGenCtx(config, tableName string) (core.Context, error) {
	ctx := core.NewGenContext()
	genConfig, err := core.ReadConfig(config)
	if err != nil {
		return ctx, err
	}
	ctx.SetGenConfig(genConfig)
	tablesConfigMap := genConfig.GetTablesConfigMap()
	if tableConfig, exist := tablesConfigMap[tableName]; exist {
		ctx.SetTableConfig(tableConfig)
	} else {
		return ctx, errors.New("can not find table [" + tableName + "] in config file")
	}
	columnsMeta, err := database.GetTableSchema(genConfig.ConnectionConfig, tableName)
	if err != nil {
		return ctx, err
	}
	ctx.SetColumnMetas(columnsMeta)
	return ctx, nil
}
