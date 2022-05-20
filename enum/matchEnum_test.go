package enum

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	m := MatchEqual
	fmt.Println(m)
	fmt.Printf("%s %d", m, m)
}
