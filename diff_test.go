package godash

import (
	"fmt"
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

func (t *Person) Show() {
	fmt.Println("woman say hi")
}

func TestDiff(t *testing.T) {
	dataList := []struct {
		dst  interface{}
		src  interface{}
		want interface{}
	}{
		{Person{Name: "Alice", Age: 20}, Person{Age: 20, Name: "Alice"}, nil},
		{Person{Name: "Alice", Age: 20}, Person{Name: "Bob", Age: 20}, Person{Name: "Bob"}},
		{Person{Name: "Alice", Age: 20}, Person{Name: "Bob", Age: 21}, Person{Name: "Bob", Age: 21}},
		{Woman{Name: "Alice", Age: 20}, Person{Name: "Alice", Age: 20}, nil},
		{Woman{Name: "Alice", Age: 20}, Person{Name: "Alice", Age: 20, Height: 170}, nil},
		{Person{Name: "Alice", Age: 20, Height: 170}, Woman{Name: "Alice", Age: 21}, Person{Age: 21}},
		{Person{Name: "Alice", Age: 20}, Woman{Name: "Bob", Age: 21}, Person{Name: "Bob", Age: 21}},
	}

	for _, c := range dataList {
		got := Diff(c.dst, c.src)
		if got != c.want {
			t.Errorf("Diff(%+v, %+v) == %+v(%T), want %+v", c.dst, c.src, got, got, c.want)
		}
	}
}
