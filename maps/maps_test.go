package maps_test

import (
	"fmt"
	"testing"
)

func TestMaps(t *testing.T) {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("set", m)
	v1 := m["k1"]
	fmt.Println("k1", v1)
	fmt.Println("k2", m["k2"])

	fmt.Println("len", len(m))
	delete(m, "k1")
	fmt.Println("map", m)

	v2, prs := m["k2"]
	fmt.Println("k2", v2)
	fmt.Println("exsit k2", prs)

	_, prs2 := m["k1"]
	fmt.Println("exist k1", prs2)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map", n)
}
