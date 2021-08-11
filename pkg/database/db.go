package database

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lizhennet/gorm-gen/pkg/core"
)

var getTableSchemaSql = "SELECT               column_name," +
	"                                         column_comment," +
	"       			        DATA_TYPE  AS column_type," +
	"IF(is_nullable = 'YES', true, false)  AS is_null," +
	" IF(column_key = 'PRI', true, false)  AS is_primary" +
	"	FROM information_schema.columns" +
	"	WHERE table_schema = ?  AND table_name = ? " +
	"ORDER BY ordinal_position"

func GetDbConnection(connectionConfig core.ConnectionConfig) (*sql.DB, error) {
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
			connectionConfig.User,
			connectionConfig.Passport,
			connectionConfig.Host,
			connectionConfig.Port,
			connectionConfig.Database))
	if err != nil {
		return nil, errors.New("connect database error,msg:" + err.Error())
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.New("connect database error,msg:" + err.Error())
	}
	return db, nil
}

func GetTableSchema(connectionConfig core.ConnectionConfig, tableName string) ([]core.ColumnMeta, error) {
	db, err := GetDbConnection(connectionConfig)
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(getTableSchemaSql, connectionConfig.Database, tableName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("table is not exist")
		}
		return nil, errors.New("query table schema error,msg:" + err.Error())
	}
	columnSchemas := make([]core.ColumnMeta, 0)
	for rows.Next() {
		column := &core.ColumnMeta{}
		err = rows.Scan(&column.ColumnName, &column.ColumnDescription, &column.ColumnType, &column.IsNullAble, &column.IsPrimary)
		if err != nil {
			return nil, errors.New("query table schema error,msg:" + err.Error())
		}
		columnSchemas = append(columnSchemas, *column)
	}
	return columnSchemas, nil
}
