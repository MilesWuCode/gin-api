package plugin

import (
	"reflect"
)

func StructToMapString(s interface{}) map[string]interface{} {
	m := make(map[string]interface{})

	v := reflect.ValueOf(s)

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		m[t.Field(i).Name] = v.Field(i).Interface()
	}

	return m
}
