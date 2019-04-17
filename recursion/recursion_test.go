package recursion

import (
	"fmt"
	"testing"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * (n - 1)
}

func TestFact(t *testing.T) {
	fmt.Println(fact(5))
}
