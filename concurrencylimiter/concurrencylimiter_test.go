package concurrencylimiter

import (
	"fmt"
	"testing"
	"time"
)

func TestDo(t *testing.T) {
	Init()
	for i := 0; i < 50; i++ {
		Do(func() error {
			time.Sleep(10)
			fmt.Println(i)
			return nil
		})
	}
}
