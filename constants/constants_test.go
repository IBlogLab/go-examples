package constants

import (
	"fmt"
	"math"
	"testing"
)

// 声明常量
const s string = "constant"

func TestConstants(t *testing.T) {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(n))
	fmt.Println(math.Sin(n))
}
