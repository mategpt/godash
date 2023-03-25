package godash

import (
	"fmt"
	"reflect"
)

func Merge(dst, src interface{}) interface{} {
	//reflect.ValueOf(&p).Elem().FieldByName(k).Set(reflect.ValueOf(v))
	//dstValue := reflect.ValueOf(&dst).Elem()x
	m := ToMap(src)
	//for k, v := range srcMap {
	//	reflect.ValueOf(&dst).Elem().FieldByName(k).Set(reflect.ValueOf(v))
	//}

	//err := mergo.Map(&dst, srcMap, mergo.WithOverride)
	//if err != nil {
	//	return err
	//}
	//return nil

	pv := reflect.ValueOf(dst)
	if pv.Kind() == reflect.Ptr {
		pv = pv.Elem()
	}
	for k, v := range m {
		//reflect.ValueOf(&p).Elem().FieldByName(k).Set(reflect.ValueOf(v))

		//fmt.Println("yyyyyyyyyyy", pv.Elem().Kind(), pv.Type())
		fmt.Println(pv)
		pv.FieldByName(k).Set(reflect.ValueOf(v))
		fmt.Printf("66666 %+v", pv) // Output: {Bob 40}
	}
	return dst
}
