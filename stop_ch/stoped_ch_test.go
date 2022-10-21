package stop_ch

import (
	"fmt"
	"testing"
	"time"
)

var (
	globalStopChan chan struct{}
	stopendChans   []chan struct{}
)

func TestStopendChan(t *testing.T) {
	globalStopChan = make(chan struct{})
	for i := 0; i < 8; i++ {
		stopendChans = append(stopendChans, make(chan struct{}))
		go func(n int) {
			defer func() {
				fmt.Println("stop: ", n)
				close(stopendChans[n])
			}()

			var timeChan <-chan time.Time
			tick := time.NewTicker(3 * time.Second)
			timeChan = tick.C

			sre := 0
			for {

				select {
				case <-timeChan:
					fmt.Println("go: ", n)
					sre++
					if sre == 2 {
						return
					}
				case <-globalStopChan:
					return
				}
			}
		}(i)
	}


	for i, v := range stopendChans {
		<-v
		fmt.Println(i, " next")
	}

	fmt.Println("all stop")

}
