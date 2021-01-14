package main

import "fmt"

func main() {
	demo()

	demo02()
}

var test string = "Test"

func demo() string {

	// 同一个代码块里面 可以使用 短变量重声明， 但是类型必须相同
	// var test string = "Test"

	test = "5"
	fmt.Println(test)

	test := 3
	fmt.Println(test)

	return "OK"
}

var block = "package"

func demo02() {
	block := "function"
	{
		block := "inner"
		fmt.Printf("The block is %s.\n", block)
	}
	fmt.Printf("The block is %s.\n", block)
}
