package syncpooler

import (
	"fmt"
	"sync"
	"testing"
)

func Test1(t *testing.T) {
	var npool sync.Pool
	fmt.Println(npool.Get())
	npool.Put("str")
	npool.Put("str1")
	npool.Put("str3")
	fmt.Println(npool.Get())
	fmt.Println(npool.Get())
	fmt.Println(npool.Get())
}
