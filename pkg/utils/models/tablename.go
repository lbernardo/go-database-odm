package models

import (
	"github.com/gertd/go-pluralize"
	"log"
	"reflect"
	"strings"
)

func GetTableNameByModel(model any) string {
	t := reflect.TypeOf(model)

	if t.Kind() != reflect.Struct && t.Kind() != reflect.Slice && t.Kind() != reflect.Ptr {
		log.Fatalln("[common/database] Was expected Struct, Slice or Ptr and we received:", t.Kind())
		return ""
	}

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
	}

	if t.Kind() == reflect.Slice {
		t = t.Elem()
	}

	field, ok := t.FieldByName("_")
	if ok {
		databaseStr := field.Tag.Get("database")
		if databaseStr != "" {
			return strings.ReplaceAll(databaseStr, "table:", "")
		}
	}
	pl := pluralize.NewClient()
	return pl.Plural(strings.ToLower(t.Name()))
}
