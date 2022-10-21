package maxHeap

import (
	"fmt"
	"testing"
)

func TestNewMaxHeap(t *testing.T) {
	st := NewMaxHeap(2)

	st.push(stat{
		name: "a",
		value: 1,
	})

	st.push(stat{
		name: "b",
		value: 2,
	})

	st.push(stat{
		name: "c",
		value: 3,
	})

	st.push(stat{
		name: "d",
		value: 4,
	})

	st.push(stat{
		name: "e",
		value: 5,
	})

	st.push(stat{
		name: "f",
		value: 6,
	})

	fmt.Println(st.get())

}
