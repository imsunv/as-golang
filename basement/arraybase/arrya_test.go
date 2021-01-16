package main

import (
	"fmt"
	"testing"
)

func TestArrayBase(t *testing.T) {
	a := make([]int, 3)
	a[0] = 1
	a[1] = 2
	a[2] = 3
	t.Log(a)
}

func TestLenAndCap(t *testing.T) {
	// 示例1。
	s1 := make([]int, 5)
	fmt.Printf("The length of s1: %d\n", len(s1))
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	fmt.Printf("The value of s1: %d\n", s1)

	s2 := make([]int, 5, 8)
	fmt.Printf("The length of s2: %d\n", len(s2))
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
	fmt.Printf("The value of s2: %d\n", s2)

	s3 := s2[3:]
	s3[0] = -1
	fmt.Printf("The length of s3: %d\n", len(s3))
	fmt.Printf("The capacity of s3: %d\n", cap(s3))
	fmt.Printf("The value of s3: %d\n", s3)

	s4 := append(s3, 1, 2)
	fmt.Println(s4[0] == s3[0])
	fmt.Printf("The length of s4: %d\n", len(s4))
	fmt.Printf("The capacity of s4: %d\n", cap(s4))
	fmt.Printf("The value of s4: %d\n", s4)

	fmt.Println("=====")
	// s3 和 s2 共用一个底层数组，因此 一个修改另一个可以看到
	fmt.Printf("The length of s2: %d\n", len(s2))
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
	fmt.Printf("The value of s2: %d\n", s2)
}
