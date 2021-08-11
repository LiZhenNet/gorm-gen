package generator

import (
	"bytes"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/lizhennet/gorm-gen/pkg/core"
	"github.com/lizhennet/gorm-gen/pkg/file"
	"github.com/lizhennet/gorm-gen/pkg/log"
)

type DalGenerator struct {
	genConfig   core.GeneratorConfig
	tableConfig core.TableConfig
	columnMetas []core.ColumnMeta
	path        string
	fileName    string
}

func (m *DalGenerator) Init(ctx core.Context) error {
	m.genConfig = ctx.Get(core.GenConfigKey).(core.GeneratorConfig)
	m.tableConfig = ctx.Get(core.TableConfigKey).(core.TableConfig)
	m.columnMetas = ctx.Get(core.ColumnMetaKey).([]core.ColumnMeta)
	m.path = core.GetPathByPackage(core.GetDalPackage(m.genConfig, m.tableConfig))
	m.fileName = strcase.ToSnake(m.tableConfig.ModelName) + "_common_dal"
	return file.CreatIsNotExist(m.path)
}

func (m *DalGenerator) Exec(ctx core.Context) error {
	modelMeta, err := genModelMeta(m.tableConfig, m.columnMetas, core.GetModelPackage(m.genConfig, m.tableConfig))
	if err != nil {
		return err
	}
	dalMeta := &core.DalMeta{ModelMeta: modelMeta}
	dalMeta.DalPackage = core.GetPackageName(core.GetDalPackage(m.genConfig, m.tableConfig))
	dalMeta.ModelModule = m.genConfig.ProjectConfig.ProjectModule + core.GetModelPackage(m.genConfig, m.tableConfig)
	temp, err := tpl.ReadFile("templates/dal.tmpl")
	if err != nil {
		return err
	}
	t, err := template.New("dal").Funcs(FuncMap()).Parse(string(temp))
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, dalMeta)
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

func (m *DalGenerator) AfterExec(ctx core.Context) error {
	log.Info("generator dal   file success, file:%s", m.fileName)
	return core.ProcessGoFile(m.path)
}
