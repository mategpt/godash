package godash

import (
	"fmt"
	"testing"

	"github.com/imdario/mergo"
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
		dst  interface{}
		src  interface{}
		want interface{}
	}{
		//{PersonM{Name: "Alice", Age: 20}, PersonM{Age: 20, Name: "Alice"}, PersonM{Name: "Alice", Age: 20}},
		{PersonM{Name: "Alice", Age: 20}, PersonM{Name: "Bob", Age: 20}, PersonM{Name: "Bob", Age: 20}},
		{PersonM{Name: "Alice", Age: 20}, PersonM{Name: "Bob", Age: 21}, PersonM{Name: "Bob", Age: 21}},
		{WomanM{Name: "Alice", Age: 20}, PersonM{Name: "Alice", Age: 20}, WomanM{Name: "Alice", Age: 20}},
		{WomanM{Name: "Alice", Age: 20}, PersonM{Name: "Alice", Age: 20, Height: 170}, WomanM{Name: "Alice", Age: 20}},
		{PersonM{Name: "Alice", Age: 20, Height: 170}, WomanM{Name: "Alice", Age: 21}, PersonM{Name: "Alice", Age: 21, Height: 170}},
		{PersonM{Name: "Alice"}, WomanM{Name: "Bob", Age: 21}, PersonM{Name: "Bob", Age: 21}},
	}

	for _, c := range dataList {
		err := mergo.Merge(&c.dst, c.src, mergo.WithOverride)
		if err != nil {
			fmt.Printf("can't merge %T to %T, %+v", c.dst, c.src, err)
		}
		if c.dst != c.want {
			t.Errorf("Merge %+v == %+v(%T), want %+v", c.src, c.dst, c.dst, c.want)
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
	p := PersonM{Name: "Alice"}
	w := PersonM{Name: "Bob", Age: 21}
	err := Merge(&p, w)
	if err != nil {
		fmt.Printf("can't merge %T to %T, %+v\n", p, w, err)
	}
	fmt.Printf("Merge %+v == %+v(%T)", w, p, p)
}
