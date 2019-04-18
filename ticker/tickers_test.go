package tickers

import (
	"fmt"
	"testing"
	"time"
)

func TestTickers(t *testing.T) {
	ticker := time.NewTicker(2 * time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(16 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
