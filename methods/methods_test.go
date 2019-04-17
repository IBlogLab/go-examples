package methods_test

import (
	"fmt"
	"testing"
)

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2 * (r.width + r.height)
}

func TestMethods(t *testing.T) {
	r := rect{width: 10, height: 5}

	fmt.Println("area", r.area())
	fmt.Println("perim", r.perim())

	rp := &r
	fmt.Println("area", rp.area())
	fmt.Println("perim", rp.perim())
}
