package core

// GeneratorConfig config for gorm-gen
type GeneratorConfig struct {
	ProjectConfig    ProjectConfig    `mapstructure:"Project"`
	ConnectionConfig ConnectionConfig `mapstructure:"Connection"`
	TablesConfig     []TableConfig    `mapstructure:"Tables"`
}

func (config GeneratorConfig) GetTablesConfigMap() map[string]TableConfig {
	tableConfigMap := make(map[string]TableConfig)
	for _, table := range config.TablesConfig {
		tableConfigMap[table.TableName] = table
	}
	return tableConfigMap
}

// ProjectConfig config for generator file
type ProjectConfig struct {
	ProjectModule string
	ModelPackage  string
	DalPackage    string
}

// ConnectionConfig config for database connection
type ConnectionConfig struct {
	Host     string
	Port     string
	User     string
	Passport string
	Database string
}

// TableConfig config for database table
type TableConfig struct {
	ModelPackage string
	DalPackage   string
	TableName    string
	ModelName    string
	Columns      []ColumnConfig
}

func (config TableConfig) GetColumnConfigMap() map[string]ColumnConfig {
	columnConfigMap := make(map[string]ColumnConfig)
	for _, column := range config.Columns {
		columnConfigMap[column.ColumnName] = column
	}
	return columnConfigMap
}

// ColumnConfig config for database table
type ColumnConfig struct {
	ColumnName string
	FieldName  string
	FieldType  string
	Tags       string
}
