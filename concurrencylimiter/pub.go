package concurrencylimiter

import (
	"fmt"
	"time"
)

type pub struct {
	ch1 chan struct{}
	ch2 chan struct{}
}

func NewPub(c1, c2 chan struct{}) *pub {
	return &pub{
		ch1: c1,
		ch2: c2,
	}
}

func (p *pub) run() {
	for  {


		time.Sleep(time.Second * 5)
		fmt.Println("pub")
		p.ch1<- struct{}{}
		<-p.ch2
	}


}
