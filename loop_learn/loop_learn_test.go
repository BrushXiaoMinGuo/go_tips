package loop_learn

import (
	"fmt"
	"testing"
)

func TestContinueLoop(t *testing.T) {

	var arr []int


loop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 2; j++ {
			if j==1 {
				continue loop
			}
			fmt.Println(i, j)
			arr = append(arr,i)
		}
	}

	fmt.Println(arr)
	arr = arr[:0]
	fmt.Println(arr)
	arr = append(arr, 0)
	fmt.Println(arr)

}
