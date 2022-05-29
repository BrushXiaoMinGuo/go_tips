package wg

import (
	"fmt"
	"sync"
	"time"
)

var (
	numWG sync.WaitGroup
)

func Init() {
	numWG.Add(1)
	go func() {
		defer numWG.Done()
		fmt.Println("start sleep")
		time.Sleep(3 * time.Second)
		fmt.Println("stop sleep")
	}()
}

func Stop() {
	numWG.Wait()
	time.Sleep(3 * time.Second)
	fmt.Println("stop function")
}
