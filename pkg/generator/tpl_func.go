package generator

import (
	"text/template"

	"github.com/lizhennet/gorm-gen/pkg/core"
)

func FuncMap() template.FuncMap {
	return map[string]interface{}{
		"concat":             func(a, b string) string { return a + b },
		"getConditionName":   func(modelMeta core.ModelMeta) string { return modelMeta.ClassName + "Condition" },
		"getUpdateModelName": func(modelMeta core.ModelMeta) string { return modelMeta.ClassName + "Updater" },
		"getDalName":         func(modelMeta core.ModelMeta) string { return modelMeta.ClassName + "CommonDal" },
		"getModelRefName":    func(modelMeta core.ModelMeta) string { return modelMeta.Package + "." + modelMeta.ClassName },
	}
}
