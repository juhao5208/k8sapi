package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/**
 * @author  巨昊
 * @date  2021/11/6 14:24
 * @version 1.15.3
 */

var db *sql.DB

func init() {
	database, err := sql.Open("mysql",
		"root:5208juhao@tcp(127.0.0.1:3306)/k8sapi2")
	if err != nil {
		panic(err)
	}
	db = database
}

func GetQueryColumns(rows *sql.Rows) ([]string, map[string]string, error) {
	columnTypes, err := rows.ColumnTypes()

	if err != nil {
		return nil, nil, err

	}

	length := len(columnTypes)

	columns := make([]string, length)

	columnTypeMap := make(map[string]string, length)

	for i, ct := range columnTypes {
		columns[i] = ct.Name()

		columnTypeMap[ct.Name()] = ct.DatabaseTypeName()

	}

	return columns, columnTypeMap, nil

}

func QueryForInterface(db *sql.DB, sqlInfo string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := db.Query(sqlInfo, args...)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	columns, columnTypeMap, err := GetQueryColumns(rows)
	if err != nil {
		return nil, err
	}
	columnLength := len(columns)
	cache := make([]interface{}, columnLength)
	for index, _ := range cache {
		var a interface{}
		cache[index] = &a
	}
	var list []map[string]interface{} //返回的切片
	for rows.Next() {
		_ = rows.Scan(cache...)
		item := make(map[string]interface{})
		for i, data := range cache {
			if ct, ok := columnTypeMap[columns[i]]; ok {
				if (ct == "VARCHAR" || ct == "DATETIME") && *data.(*interface{}) != nil {
					item[columns[i]] = string((*data.(*interface{})).([]byte))
				} else {
					item[columns[i]] = *data.(*interface{})
				}
			} else {
				item[columns[i]] = *data.(*interface{})
			}
		}
		list = append(list, item)
	}
	return list, nil
}

func main() {
	/*server := goft.Ignite().Mount("",
		controllers.NewUserCtl()).Attach()
	server.Launch()*/

	li, err := QueryForInterface(db, "select * from user")
	if err != nil {
		panic(err)
	}
	fmt.Println(li)

}
