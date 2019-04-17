package channels

import (
	"fmt"
	"testing"
)

func TestChannelBuffering(t *testing.T) {
	message := make(chan string, 2)
	message <- "buffer"
	message <- "channel"

	fmt.Println(<-message)
	fmt.Println(<-message)
	message <- "buffer2"
	fmt.Println(<-message)
}
