package godash

import (
	"fmt"
	"github.com/imdario/mergo"
	"reflect"
	"testing"
)

type PersonM struct {
	Name   string
	Age    int
	Height int
}
type WomanM struct {
	Name string
	Age  int
}

func TestDiffM(t *testing.T) {
	dataList := []struct {
		dst  PersonM
		src  interface{}
		want interface{}
	}{
		{PersonM{Name: "Alice", Age: 20}, PersonM{Age: 20, Name: "Alice"}, PersonM{Name: "Alice", Age: 20}},
		{PersonM{Name: "Alice", Age: 20}, PersonM{Name: "Bob", Age: 20}, PersonM{Name: "Bob", Age: 20}},
		{PersonM{Name: "Alice", Age: 20}, PersonM{Name: "Bob", Age: 21}, PersonM{Name: "Bob", Age: 21}},
		//{&WomanM{Name: "Alice", Age: 20}, PersonM{Name: "Alice", Age: 20}, WomanM{Name: "Alice", Age: 20}},
		//{&WomanM{Name: "Alice", Age: 20}, PersonM{Name: "Alice", Age: 20, Height: 170}, WomanM{Name: "Alice", Age: 20}},
		{PersonM{Name: "Alice", Age: 20, Height: 170}, WomanM{Name: "Alice", Age: 21}, PersonM{Name: "Alice", Age: 21, Height: 170}},
		{PersonM{Name: "Alice"}, WomanM{Name: "Bob", Age: 21}, PersonM{Name: "Bob", Age: 21}},
	}

	for _, c := range dataList {
		srcMap := Diff(c.dst, c.src)
		//Merge(&c.dst, c.src)
		//err := Merge(&c.dst, c.src)
		//if err != nil {
		//	fmt.Printf("can't merge %T to %T, %+v", c.dst, c.src, err)
		//}
		dstType := reflect.ValueOf(&c.dst).Elem().Kind()
		fmt.Println("000000000000000", dstType, c.dst)
		err := mergo.Map(&c.dst, srcMap, mergo.WithOverride)
		fmt.Println("111111111111111", dstType, c.dst)
		//fmt.Printf("Merge %+v(%T) == %+v(%T), want %+v(%T)", c.src, c.src, c.dst, c.dst, c.want, c.want)
		if err != nil {
			fmt.Printf("can't merge %T to %T, %+v", c.dst, c.src, err)
		}
		if !reflect.DeepEqual(c.dst, c.want) {
			t.Errorf("Merge %+v(%T) == %+v(%T), want %+v(%T)", c.src, c.src, c.dst, c.dst, c.want, c.want)
		}
	}

	//
	//for _, c := range dataList {
	//	err := Merge(&c.dst, c.src)
	//	if err != nil {
	//		fmt.Printf("can't merge %T to %T, %+v\n", c.dst, c.src, err)
	//	}
	//	if c.dst != c.want {
	//		t.Errorf("Merge %+v == %+v(%T), want %+v", c.src, c.dst, c.dst, c.want)
	//	}
	//}
	//p := PersonM{Name: "Alice"}
	//w := PersonM{Name: "Bob", Age: 21}
	//err := Merge(&p, w)
	//if err != nil {
	//	fmt.Printf("can't merge %T to %T, %+v\n", p, w, err)
	//}
	//fmt.Printf("Merge %+v == %+v(%T)", w, p, p)
}
