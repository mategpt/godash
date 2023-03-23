package godash

import (
	"reflect"
)

func Diff(structOld, structNew interface{}) interface{} {
	dst := reflect.ValueOf(structOld)
	src := reflect.ValueOf(structNew)
	if src.Kind() == reflect.Ptr {
		src = src.Elem()
	}
	if dst.Kind() == reflect.Ptr {
		dst = dst.Elem()
	}

	diff := reflect.New(dst.Type()).Elem()
	for i := 0; i < src.NumField(); i++ {
		f1 := src.Field(i)
		if i < dst.NumField() {
			f2 := dst.Field(i)
			if !reflect.DeepEqual(f1.Interface(), f2.Interface()) {
				diff.Field(i).Set(f1)
			}
		}
	}
	if reflect.DeepEqual(diff.Interface(), reflect.Zero(diff.Type()).Interface()) {
		return nil
	}
	return diff.Interface()
}
