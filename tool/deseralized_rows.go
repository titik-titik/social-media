package tool

import (
	"database/sql"
	"reflect"
)

func DeserializeRows(rows *sql.Rows, entityType interface{}) interface{} {
	var results []interface{}

	for rows.Next() {
		entity := reflect.New(reflect.TypeOf(entityType)).Interface()
		scanErr := rows.Scan(reflect.ValueOf(entity).Elem().Addr().Interface())

		if scanErr != nil {
			panic(scanErr)
		}
		results = append(results, entity)
	}
	return results
}
