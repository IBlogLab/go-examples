package channels

import (
	"fmt"
	"testing"
)

/* pings is send message, can't receive value */
func ping(pings chan<- string, msg string) {
	pings <- msg
}

/* pings is receive message, can't send message;
* pongs is send message, can't receive message
 */
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func TestChannelDirections(t *testing.T) {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "send message...")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
