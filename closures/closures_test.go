package closures

import (
	"fmt"
	"testing"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func TestClosures(t *testing.T) {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInt1 := intSeq()
	fmt.Println(nextInt1())
	fmt.Println(nextInt1())
	fmt.Println(nextInt1())
}
