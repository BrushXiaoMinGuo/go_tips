package concurrencylimiter

import "fmt"

const (
	MaxConcurrencyLimit = 2
	MaxDuration         = 5
)

var ch chan struct{}

func Init() {
	ch = make(chan struct{}, MaxConcurrencyLimit)
}

func Do(f func() error) error {
	select {
	case ch <- struct{}{}:
		fmt.Println("not enough")
		err := f()
		<-ch
		return err
	default:
	}

	t := Get(MaxDuration)
	select {
	case ch <- struct{}{}:
		Put(t)
		fmt.Println("has empty")
		err := f()
		<-ch
		return err
	case <-t.C:
		Put(t)
		return nil
	}
}
