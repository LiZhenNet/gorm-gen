package core

import (
	"path/filepath"
	"strings"
)

func GetStringWithDefault(value, defaultValue string) string {
	if value != "" {
		return value
	}
	return defaultValue
}

func GetPackageName(path string) string {
	return filepath.Base(path)
}

func GetPathByPackage(p string) string {
	if strings.HasPrefix(p, "/") {
		return "." + p
	}
	return "./" + p
}

func GetModelPackage(genConfig GeneratorConfig, tableConfig TableConfig) string {
	return GetStringWithDefault(tableConfig.ModelPackage, genConfig.ProjectConfig.ModelPackage)
}

func GetDalPackage(genConfig GeneratorConfig, tableConfig TableConfig) string {
	return GetStringWithDefault(tableConfig.DalPackage, genConfig.ProjectConfig.DalPackage)
}
