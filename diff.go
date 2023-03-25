package godash

import (
	"reflect"
)

func DiffRetStruct(structOld, structNew interface{}) interface{} {
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

func Diff(structOld, structNew interface{}) map[string]interface{} {
	dst := reflect.ValueOf(structOld)
	src := reflect.ValueOf(structNew)
	srcType := reflect.TypeOf(structNew)

	if src.Kind() == reflect.Ptr {
		src = src.Elem()
	}
	if srcType.Kind() == reflect.Ptr {
		srcType = srcType.Elem()
	}
	if dst.Kind() == reflect.Ptr {
		dst = dst.Elem()
	}

	diffMap := map[string]interface{}{}
	for i := 0; i < src.NumField(); i++ {
		key := srcType.Field(i).Name
		value := src.Field(i).Interface()

		dstField := dst.FieldByName(key)
		if dstField.IsValid() {
			dstValue := dst.FieldByName(key).Interface()
			if !reflect.DeepEqual(value, dstValue) {
				diffMap[key] = value
			}
		}
	}
	if len(diffMap) == 0 {
		return nil
	}
	return diffMap
}
