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
	fun2 := fibonacci()

	PrintFibonacci(fun)
	PrintFibonacci(fun2)
}

func inc(x int) func(y int) int {
	return func(y int) int {
		return x + y
	}
}

func Test_Function_Currying(t *testing.T) {
	inc2 := inc(2)

	fmt.Println(inc2(5))
}

func evenFilter(num int) bool {
	return (num & 1) == 0
}

func multiplyByThree(num int) int {
	return num * 3
}

func convertString(num int) string {
	return fmt.Sprintf("The Number: %d", num)
}

func Test_process(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, n := range nums {
		if evenFilter(n) {
			fmt.Println(convertString(multiplyByThree(n)))
		}
	}
}
