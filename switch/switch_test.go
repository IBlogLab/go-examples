package switch_test

import (
	"fmt"
	"testing"
	"time"
)

func TestSwitch(t *testing.T) {
	i := 2
	fmt.Print("Write", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("it's Weekday")
	}

	t1 := time.Now()
	switch {
	case t1.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("str")
}
