package wg

import (
	"fmt"
	"testing"
)

func TestWG(t *testing.T) {

	Init()
	Stop()
	fmt.Println("test stop")
}
