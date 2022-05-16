package selecter

import (
	"fmt"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	var ch chan bool
	ch = make(chan bool,1)

	tt := time.NewTimer(2*time.Second)
	fmt.Println(time.Now())
	ch<-false

	go func() {
		time.Sleep(1*time.Second)
		<-ch
	}()

	select {
	case ch<-false:
		fmt.Println("ok")
		fmt.Println(time.Now())
	case <-tt.C:
		fmt.Println(time.Now())

	}
}

func Test2(t *testing.T) {
	var ch chan bool
	ch = make(chan bool,1)

	tt := time.NewTimer(2*time.Second)
	fmt.Println(time.Now())
	ch<-false

	go func() {
		time.Sleep(1*time.Second)
		<-ch
	}()

	select {
	case ch<-false:
		fmt.Println("ok")
		fmt.Println(time.Now())
	case <-tt.C:
		fmt.Println(time.Now())
	default:
		fmt.Println(time.Now())

	}
}