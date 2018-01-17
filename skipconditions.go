package gopartial

import (
	"reflect"
	"strings"

	"github.com/roserocket/roserocket/utils"
)

// SkipReadOnly skips all field that has tag readonly
func SkipReadOnly(field reflect.StructField) bool {
	props := strings.Split(field.Tag.Get("props"), ",")
	return utils.IndexOf(props, "readonly") >= 0
}

// SkipConditions collection of all skip conditions
var SkipConditions = []func(reflect.StructField) bool{
	SkipReadOnly,
}
