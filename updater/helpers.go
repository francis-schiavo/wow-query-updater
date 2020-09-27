package updater

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"wow-query-updater/connections"
	"wow-query-updater/datasets"
)

func insertOnce(data interface{})  {
	db := connections.GetDBConn()

	_, err := db.Model(data).
		OnConflict("(id) DO NOTHING").
		Insert()
	if err != nil {
		handleError(err, data)
	}
}

func insertOnceUpdate(data interface{}, fields ...string)  {
	db := connections.GetDBConn()

	var expression string
	for _, field := range fields {
		expression += fmt.Sprintf("%s = EXCLUDED.%s,", field, field)
	}
	expression = strings.Trim(expression, ",")

	_, err := db.Model(data).
		OnConflict("(id) DO UPDATE").
		Set(expression).
		Insert()
	if err != nil {
		handleError(err, data)
	}
}

func insertOnceExpr(data interface{}, conflictExpr string, fields ...string)  {
	db := connections.GetDBConn()

	expression := ""
	for _, field := range fields {
		expression += fmt.Sprintf("%s = EXCLUDED.%s,", field, field)
	}
	expression = strings.Trim(expression, ",")

	_, err := db.Model(data).OnConflict(conflictExpr).Set(expression).Insert()
	if err != nil {
		handleError(err, data)
	}
}

func insert(data interface{})  {
	db := connections.GetDBConn()
	_, err := db.Model(data).OnConflict("DO NOTHING").Insert()
	if err != nil {
		handleError(err, data)
	}
}

func handleError(err error, data interface{}) {
	name := reflect.TypeOf(data).Elem().Name()
	//log.Fatalf("%s: Error: %v Data: %v", name, err, data)

	db := connections.GetDBConn()
	log.Printf("%s: Error: %v Data: %v", name, err, data)
	db.Model(&datasets.UpdateError{
		Error: fmt.Sprintf("%v %v", err, data),
	}).Insert()
}