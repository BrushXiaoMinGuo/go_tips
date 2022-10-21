package _chan

import (
	"testing"
)

func TestMakeChan(t *testing.T) {
	var noBufferChan chan int
	var bufferChan chan int
	noBufferChan = make(chan int)
	bufferChan = make(chan int, 10)

	for i := 0; i < 10; i++ {
		go func(n int) {
			noBufferChan<-n
		}(i)
	}

	for i := 0; i < 10; i++ {
		bufferChan <- i
	}

}
