package bitmap

import (
	"bytes"
	"fmt"
	"github.com/pilosa/pilosa/v2/roaring"
	"testing"
)

func TestNewBitMap(t *testing.T) {
	b := roaring.NewBTreeBitmap()
	b.Add(1,2,3)
	b.Add(3,4,5)


	fmt.Println(b.Contains(15))
	bb := &bytes.Buffer{}
	fmt.Println(fmt.Sprint(bb))

	a := roaring.NewBTreeBitmap()
	a.Add(6,7,5)

	c := b.Union(a)
	fmt.Println(c.Contains(7))
}
