/**
 * @Desc
 * @author zjhfyq
 * @data 2018/3/20 11:06.
 */
package goGenerater

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


//获取数据库连接指针
func GetDB(ds DataSource) (db *sql.DB) {
	/*DSN数据源名称
	[username[:password]@][protocol[(address)]]/dbname[?param1=value1¶mN=valueN]
	user@unix(/path/to/socket)/dbname
	user:password@tcp(localhost:5555)/dbname?charset=utf8&autocommit=true
	user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname?charset=utf8mb4,utf8
	user:password@/dbname
	无数据库: user:password@/
	*/
	db, err := sql.Open(ds.DriverName, ds.DataSourceName)
	if err != nil {
		panic(err)
	}
	return
}

/*
select table_name from information_schema.tables where table_schema='go-generater' and table_type='base table';
*/
func GetTableNames(db *sql.DB, dbName string) (tableNames []string) {
	sql := "select table_name from information_schema.tables" +
		" where table_schema='" + dbName + "' and " +
		"table_type='base table'"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	} else {
		for rows.Next() {
			var tableName string
			err := rows.Scan(&tableName)
			if err != nil {
				panic(err)
			} else {
				tableNames = append(tableNames, tableName)
			}
		}
	}
	return tableNames
}

/**
SHOW FULL COLUMNS FROM db_name.table_name
*/
func GetFieldByTableName(db *sql.DB, dbName string, tableName string) (fields []Field) {
	sql := "SHOW FULL COLUMNS FROM "+dbName+"."+tableName
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	} else {
		for rows.Next() {
			var field Field = Field{}
			err := rows.Scan(&field.FieldName, &field.Type, &field.Collation,
				&field.Null, &field.Key, &field.Default,
				&field.Extra, &field.Privileges, &field.Comment)
			if err != nil {
				panic(err)
			} else {
				fields = append(fields, field)
			}
		}
	}
	return
}
