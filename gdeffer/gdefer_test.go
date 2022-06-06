package gdeffer

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {

	for i := 3; i > 0; i-- {
		defer fmt.Println(i)
	}

	for i := 7; i > 4 ;i-- {
		 defer func() {
		 	fmt.Println(i)
		 }()
	}

	for i:=10; i>7;i--{
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}

}
