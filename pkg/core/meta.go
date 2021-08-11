package core

type ColumnMeta struct {
	ColumnName        string
	ColumnDescription string
	ColumnType        string
	IsNullAble        bool
	IsPrimary         bool
}

type FieldMeta struct {
	FieldName        string
	FieldType        string
	FiledTags        []string
	FieldDescription string
	ColumnName       string
	IsNullAble       bool
}

type ModelMeta struct {
	Package   string
	Imports   []string
	ClassName string
	TableName string
	Fields    []FieldMeta
}

type DalMeta struct {
	ModelMeta
	DalPackage  string
	ModelModule string
}
