package gopartial

import (
	"reflect"
	"strings"
)

const readOnlyTag = "readonly"

// SkipReadOnly skips all field that has tag readonly
func SkipReadOnly(field reflect.StructField) bool {
	props := strings.Split(field.Tag.Get("props"), ",")

	for _, v := range props {
		if v == readOnlyTag {
			return true
		}
	}

	return false
}

// SkipConditions collection of all skip conditions
var SkipConditions = []func(reflect.StructField) bool{
	SkipReadOnly,
}
