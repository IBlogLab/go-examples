package channels

import (
	"fmt"
	"testing"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func TestChannelSync(t *testing.T) {
	done := make(chan bool, 1)
	go worker(done)

	res := <-done
	fmt.Print(res)
}
