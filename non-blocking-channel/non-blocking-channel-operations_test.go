package non_blocking_channel

import (
	"fmt"
	"testing"
)

func TestNonBlocking(t *testing.T) {
	messages := make(chan string)
	signal := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received msg:", msg)
	default:
		fmt.Println("non message")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("send msg", msg)
	default:
		fmt.Println("no send msg")
	}

	select {
	case msg := <-messages:
		fmt.Println("received msg:", msg)
	case sig := <-signal:
		fmt.Println("received sign", sig)
	default:
		fmt.Println("non activity")
	}
	// signal <- true
}
