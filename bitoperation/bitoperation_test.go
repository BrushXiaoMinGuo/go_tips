package bitoperation

import (
	"fmt"
	"testing"
)

//留神看这里 取模的操作应用的是&而不是% 这是因为位运算(&)效率要比取模运算(%)高很多，次要起因是位运算间接对内存数据进行操作，不须要转成十进制，因而处理速度十分快
//
// a % b == a & (b - 1) 前提：b 为 2^n

func TestFastMod(t *testing.T) {
	size := 1 << 10
	sizeNum := 1024
	a := 2023
	fmt.Println(a % sizeNum)
	fmt.Println(a & (size - 1))
	fmt.Println(a % size)
}

func TestFastRem(t *testing.T) {
	size := 1 << 10
	size = size -1
	sizeNum := 1023
	a := 2023
	fmt.Println(a % sizeNum)
	fmt.Println(a & (size - 1))
	fmt.Println(a % size)
}