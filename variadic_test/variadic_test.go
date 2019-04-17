package variadic_test

import (
	"fmt"
	"testing"
)

func variadic(nums ...int) {
	fmt.Println(nums, " ")
	var sum = 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum", sum)
}

func TestVariadic(t *testing.T) {
	variadic(1, 2, 3)
	variadic(1, 3, 5)

	nums := []int{1, 2, 3, 4, 5}
	variadic(nums...)
}
