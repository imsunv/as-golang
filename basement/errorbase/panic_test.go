package errorbase

import (
	"fmt"
	"testing"
)

func caller1() {
	fmt.Println("Enter function caller1.")
	caller2()
	fmt.Println("Exit function caller1.")
}

func caller2() {
	fmt.Println("Enter function caller2.")
	s1 := []int{0, 1, 2, 3, 4}
	e5 := s1[5]
	_ = e5
	fmt.Println("Exit function caller2.")
}

func caller3() {
	defer func() {
		fmt.Println("Enter defer function.")
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}
		fmt.Println("Exit defer function.")
	}()
	caller1()
}

func Test_C1(t *testing.T) {
	fmt.Println("Enter function main.")
	caller1()
	fmt.Println("Exit function main.")
}

func Test_C2(t *testing.T) {
	fmt.Println("Enter function main.")
	caller2()
	fmt.Println("Exit function main.")
}

func Test_C3(t *testing.T) {
	fmt.Println("Enter function main.")
	caller3()
	fmt.Println("Exit function main.")
}
