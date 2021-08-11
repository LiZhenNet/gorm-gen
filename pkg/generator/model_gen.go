package generator

import (
	"bytes"
	"embed"
	"fmt"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/lizhennet/gorm-gen/pkg/core"
	"github.com/lizhennet/gorm-gen/pkg/file"
	"github.com/lizhennet/gorm-gen/pkg/log"
)

//go:embed templates
var tpl embed.FS

type ModelGenerator struct {
	genConfig   core.GeneratorConfig
	tableConfig core.TableConfig
	columnMetas []core.ColumnMeta
	path        string
	fileName    string
}

func (m *ModelGenerator) Init(ctx core.Context) error {
	m.genConfig = ctx.Get(core.GenConfigKey).(core.GeneratorConfig)
	m.tableConfig = ctx.Get(core.TableConfigKey).(core.TableConfig)
	m.columnMetas = ctx.Get(core.ColumnMetaKey).([]core.ColumnMeta)
	m.path = core.GetPathByPackage(core.GetModelPackage(m.genConfig, m.tableConfig))
	m.fileName = strcase.ToSnake(m.tableConfig.ModelName)
	return file.CreatIsNotExist(m.path)
}

func (m *ModelGenerator) Exec(ctx core.Context) error {
	modelMeta, err := genModelMeta(m.tableConfig, m.columnMetas, core.GetModelPackage(m.genConfig, m.tableConfig))
	if err != nil {
		return err
	}
	temp, err := tpl.ReadFile("templates/model.tmpl")
	if err != nil {
		return err
	}
	t, err := template.New("model").Funcs(FuncMap()).Parse(string(temp))
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, modelMeta)
	if err != nil {
		return err
	}
	err = file.WriteFile(m.path, m.fileName, buf)
	if err != nil {
		return err
	}
	_ = core.ProcessGoFile(m.path)
	return nil
}

func (m *ModelGenerator) AfterExec(ctx core.Context) error {
	log.Info("generator model file success, file:%s", m.fileName)
	return core.ProcessGoFile(m.path)
}

func genModelMeta(tableConfig core.TableConfig, columnMeta []core.ColumnMeta, pkg string) (core.ModelMeta, error) {
	modelMeta := &core.ModelMeta{
		Package:   core.GetPackageName(pkg),
		ClassName: tableConfig.ModelName,
		TableName: tableConfig.TableName,
	}
	columnConfigMap := tableConfig.GetColumnConfigMap()
	fieldMetas := make([]core.FieldMeta, 0)
	for _, column := range columnMeta {
		fieldType, err := core.GetGoType(column.ColumnType)
		if err != nil {
			return core.ModelMeta{}, err
		}
		fieldMeta := core.FieldMeta{
			ColumnName:       column.ColumnName,
			FieldName:        strcase.ToCamel(column.ColumnName),
			FieldType:        fieldType,
			FieldDescription: column.ColumnDescription,
			FiledTags:        []string{fmt.Sprintf(core.GormTag, column.ColumnName), fmt.Sprintf(core.JsonTag, column.ColumnName)},
			IsNullAble:       column.IsNullAble,
		}
		columnConfig := columnConfigMap[column.ColumnName]
		overrideFieldMetaByConfig(&fieldMeta, columnConfig)
		fieldMetas = append(fieldMetas, fieldMeta)
	}
	modelMeta.Fields = fieldMetas
	return *modelMeta, nil
}

func overrideFieldMetaByConfig(fieldMeta *core.FieldMeta, columnConfig core.ColumnConfig) {
	fieldMeta.FieldName = core.GetStringWithDefault(columnConfig.FieldName, fieldMeta.FieldName)
	fieldMeta.FieldType = core.GetStringWithDefault(columnConfig.FieldType, fieldMeta.FieldType)
	if columnConfig.Tags != "" {
		fieldMeta.FiledTags = append(fieldMeta.FiledTags, columnConfig.Tags)
	}
}
