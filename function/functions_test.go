package function_test

import (
	"fmt"
	"testing"
)

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, string) {
	return 3, "7"
}

func TestFunc(t *testing.T) {
	res := plus(1, 2)
	fmt.Println("plus", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("plusPlus", res)

	i, s := vals()
	fmt.Println(i, s)
}
