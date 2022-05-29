package stop_ch

import (
	"fmt"
	"testing"
	"time"
)

var (
	stopCH <-chan struct{}
	timeCH <-chan time.Time
)

func TestStopCH(t *testing.T) {

	stopCH := make(chan struct{})
	ticker := time.NewTicker(3 * time.Second)
	timeCH := ticker.C

	go func() {
		for {
			select {
			case <-timeCH:
				fmt.Println("time ok")
			case <-stopCH:
				fmt.Println("stop")
				return
			}
		}
	}()

	time.Sleep(9 * time.Second)
	fmt.Println("main stop 1")
	close(stopCH)
	fmt.Println("main stop 2")
}
