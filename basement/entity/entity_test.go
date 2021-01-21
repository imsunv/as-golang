package main

import (
	"fmt"
	"testing"
)

func Test_Interface_Nil(t *testing.T) {

	// 我们只声明而不初始化，或者显式地赋给它nil，否则接口变量的值就不会为nil。
	check := func(in interface{}) {
		if in == nil {
			fmt.Println("input is nil")
		} else {
			fmt.Printf("input not nil, value: %v \n", in)
		}
	}

	var a map[string]int
	var b int
	var c string
	var d []int
	var e func()

	// Nil
	check(nil)

	// Not Nil
	check(a)
	check(b)
	check(c)
	check(d)
	check(e)

	check(1)
	check("aaa")
	check([]int{1, 2, 3})
	check([]string{"a", "b", "c"})
	check(false)
	check(func() {})
}

type User struct {
	Name string
	Age  int
}

// *User 和 User 的区别是什么
func (user *User) id() string {
	return user.Name
}

func Test_Type_Converter(t *testing.T) {

	check := func(input interface{}) {
		if v, ok := input.(int); ok {
			fmt.Printf("int: %d \n", v)
		} else if id, ok := input.(Id); ok {
			fmt.Printf("ID: %s \n", id.id())
		} else {
			fmt.Printf("unknown type: %v \n", input)
		}
	}

	check(1)
	check("aaa")
	check([]int{1, 2, 3})
	check([]string{"a", "b", "c"})
	check(false)
	check(func() {})
	check(User{"user", 10})
	check(&User{"user", 10})
}

type Id interface {
	id() string
}
