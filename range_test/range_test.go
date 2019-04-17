package range_test

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	a := []int{1, 2, 3}
	sum := 0
	for _, num := range a {
		sum += num
	}
	fmt.Println("sum", sum)

	for i, num := range a {
		fmt.Println("index", i)
		fmt.Println("v", num)
	}

	m := map[string]string{"far": "1", "bar": "2"}
	for k, v := range m {
		fmt.Println("k", k, "v", v)
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
