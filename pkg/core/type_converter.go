package core

import "errors"

func GetGoType(dbType string) (string, error) {
	switch dbType {
	case "smallint", "tinyint":
		return "int16", nil
	case "int":
		return "int32", nil
	case "bigint":
		return "int64", nil
	case "char", "varchar", "text", "mediumtext", "longtext":
		return "string", nil
	case "float", "double", "decimal":
		return "float64", nil
	case "bit":
		return "uint64", nil
	case "date", "datetime", "timestamp":
		return "time.Time", nil
	default:
		return "", errors.New("not support db type:" + dbType)
	}
}
