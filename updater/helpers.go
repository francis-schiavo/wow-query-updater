package updater

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"log"
	"reflect"
	"strconv"
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

func asInt(in string) int {
	result, _ := strconv.Atoi(in)
	return result
}

func handleError(error error, data interface{}) {
	db := connections.GetDBConn()
	name := reflect.TypeOf(data).Elem().Name()
	log.Printf("%s: Error: %v Data: %v", name, error, data)
	errType := reflect.TypeOf(error)

	if errType == reflect.TypeOf("") {
		db.Model(&datasets.UpdateError{
			Endpoint: "Generic",
			Error:    error.Error(),
		}).Insert()
	} else {
		err := error.(pg.Error)

		if err.Field(67) == "23503" {
			recordID := 0
			_, found := reflect.TypeOf(data).Elem().FieldByName("ID")
			if found {
				recordID = int(reflect.ValueOf(data).Elem().FieldByName("ID").Int())
			}

			db.Model(&datasets.UpdateError{
				Endpoint: err.Field(116),
				RecordID: recordID,
				Error:    err.Field(68),
			}).Insert()
		} else if err.Field(67) == "42P10" {
			db.Model(&datasets.UpdateError{
				Endpoint: name,
				Error:    fmt.Sprintf("%v %v", err, data),
			}).Insert()
		} else {
			db.Model(&datasets.UpdateError{
				Error: fmt.Sprintf("%v %v", err, data),
			}).Insert()
		}
	}
}