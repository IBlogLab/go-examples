package values

import (
	"fmt"
	"testing"
)

func TestValues(t *testing.T) {
	// 字符串能够使用 '+' 连接
	fmt.Println("go" + "lang")

	// 整型和浮点型
	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0/3.0=", 7.0/3.0)

	// 布尔类型
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
