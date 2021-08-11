package core

import (
	"context"
)

type Context struct {
	context.Context
	data map[string]interface{}
}

func NewGenContext() Context {
	return Context{
		Context: context.Background(),
		data:    map[string]interface{}{},
	}
}
func (c Context) Set(key string, value interface{}) {
	c.data[key] = value
}

func (c Context) Get(key string) interface{} {
	if v, exist := c.data[key]; exist {
		return v
	}
	return nil
}

func (c Context) GetGenConfig() GeneratorConfig {
	return c.Get(GenConfigKey).(GeneratorConfig)
}
func (c Context) SetGenConfig(config GeneratorConfig) {
	c.Set(GenConfigKey, config)
}

func (c Context) GetTableConfig() TableConfig {
	return c.Get(TableConfigKey).(TableConfig)
}
func (c Context) SetTableConfig(config TableConfig) {
	c.Set(TableConfigKey, config)
}

func (c Context) GetColumnMetas() []ColumnMeta {
	return c.Get(ColumnMetaKey).([]ColumnMeta)
}
func (c Context) SetColumnMetas(v []ColumnMeta) {
	c.Set(ColumnMetaKey, v)
}
