package structs

import (
	"fmt"
	"testing"
)

type person struct {
	name string
	age  int
}

func TestStructs(t *testing.T) {
	fmt.Println(person{"lance", 20})
	fmt.Println(person{name: "lance", age: 27})
	fmt.Println(person{name: "lance"})
	fmt.Println(&person{name: "lance", age: 27})

	s := person{name: "bob", age: 40}
	fmt.Println(s.name, s.age)

	sp := &s
	fmt.Println(sp.age)
	sp.age = 50
	fmt.Println(sp.age)
	fmt.Println(s.age)
}
