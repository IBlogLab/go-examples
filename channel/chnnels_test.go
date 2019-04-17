package channels

import (
	"fmt"
	"testing"
)

func TestChannels(t *testing.T) {
	msgs := make(chan string)
	go func() {
		msgs <- "ping"
	}()

	msg := <-msgs
	fmt.Println(msg)
}
