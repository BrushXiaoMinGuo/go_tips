package concurrencylimiter

import (
	"fmt"
	"sync"
	"time"
)

var timerPool sync.Pool

func Get(d time.Duration) *time.Timer {
	if v := timerPool.Get(); v != nil {
		t := v.(*time.Timer)
		if t.Reset(d) {
			fmt.Println("error")
		}
		return t
	}
	return time.NewTimer(d)
}

func Put(t *time.Timer) {
	if !t.Stop() {
		select {
		case <-t.C:
		default:

		}
	}
	timerPool.Put(t)
}
