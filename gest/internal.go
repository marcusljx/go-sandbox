package gest

import (
	"reflect"
	"unsafe"
)

// extractUnexportedField takes
func extractUnexportedField(fieldName string, v reflect.Value) reflect.Value {
	field := v.FieldByName(fieldName)
	result := reflect.NewAt(
		field.Type(),
		unsafe.Pointer(field.UnsafeAddr()),
	).Elem()
	return result
}
