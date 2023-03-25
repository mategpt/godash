package godash

import (
	"testing"
)

type Person struct {
	Name   string
	Age    int
	Height int
}
type Woman struct {
	Name string
	Age  int
}

func TestDiff(t *testing.T) {
	dataList := []struct {
		dst  interface{}
		src  interface{}
		want map[string]interface{}
	}{
		{Person{Name: "Alice", Age: 20}, &Person{Age: 20, Name: "Alice"}, nil},
		{Person{Name: "Alice", Age: 20}, Person{Name: "Bob", Age: 20}, map[string]interface{}{"Name": "Bob"}},
		{Person{Name: "Alice", Age: 20}, Person{Name: "Bob", Age: 21}, map[string]interface{}{"Name": "Bob", "Age": 21}},
		{Woman{Name: "Alice", Age: 20}, Person{Name: "Alice", Age: 20}, nil},
		{Woman{Name: "Alice", Age: 20}, Person{Name: "Alice", Age: 20, Height: 170}, nil},
		{Person{Name: "Alice", Age: 20, Height: 170}, Woman{Name: "Alice", Age: 21}, map[string]interface{}{"Age": 21}},
		{Person{Name: "Alice", Age: 20}, Woman{Name: "Bob", Age: 21}, map[string]interface{}{"Name": "Bob", "Age": 21}},
	}

	for _, c := range dataList {
		got := Diff(c.dst, c.src)
		if !IsEqual(got, c.want) {
			t.Errorf("Diff(%+v, %+v) == %+v(%T), want %+v(%T)", c.dst, c.src, got, got, c.want, c.want)
		}
	}
}
