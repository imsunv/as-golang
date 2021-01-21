package function

import (
	"fmt"
	"testing"
)

// 对象是附有行为的数据，而闭包是附有数据的行为

func TestAddr(t *testing.T) {

	addr := Addr()

	fmt.Println(addr(10))
	fmt.Println(addr(10))
	fmt.Println(addr(10))
}

func TestFibonacci(t *testing.T) {
	fun := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d , ", fun())
	}
}

func TestFibonacci2(t *testing.T) {
	fun := fibonacci()

	PrintFibonacci(fun)
}
