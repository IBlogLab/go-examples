package for_test

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T) {
	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}

	for i := 7; i <= 9; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("for loop")
		break
	}

	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
