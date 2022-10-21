package concurrencylimiter

import (
	"fmt"
	"time"
)

type sub struct {
	ch1 chan struct{}
	ch2 chan struct{}
}

func NewSub(c1 , c2 chan struct{}) *sub {
	return &sub{
		ch1: c1,
		ch2: c2,
	}
}

func (s *sub) run() {
	for  {

		<-s.ch1
		time.Sleep(time.Second * 5)
		fmt.Println("sub")

		s.ch2<- struct{}{}
	}




}
