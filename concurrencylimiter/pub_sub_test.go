package concurrencylimiter

import (
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestSub(t *testing.T) {

	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)
	p := NewPub(ch1, ch2)
	s := NewSub(ch1, ch2)



		go s.run()
	time.Sleep(time.Second*2)
	go p.run()

	done := make(chan os.Signal)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

	select {
	case <-done:
	}




}
