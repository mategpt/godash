package godash

import (
	"reflect"
)

func Merge(dst, src interface{}) interface{} {
	srcMap := ToMap(src)
	dstValue := reflect.ValueOf(dst).Elem()
	for k, v := range srcMap {
		dstField := dstValue.FieldByName(k)
		dstField.Set(reflect.ValueOf(v))
	}
	return dst
}
