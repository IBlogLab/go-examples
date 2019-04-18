package closingchannels

import (
	"fmt"
	"testing"
)

func TestClosingChannels(t *testing.T) {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs

			if more {
				fmt.Println("received job:", j)
			} else {
				fmt.Println("finished to receive all jobs.")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send job:", j)
	}

	close(jobs)
	fmt.Println("send all jobs")
	<-done
}
