package pointers

import (
	"fmt"
	"testing"
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func TestPointer(t *testing.T) {
	i := 1
	fmt.Println("init", i)

	zeroval(i)
	fmt.Println("zeroval", i)

	zeroptr(&i)
	fmt.Println("zeroptr", i)

	fmt.Println("point", &i)
}
