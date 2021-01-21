package main

import (
	"fmt"
	"unsafe"
)

type test interface {
}

func main() {
	m := make(map[string]string)

	c := func(key string) string {
		return key + " _aa"
	}

	//l := []int{1, 2, 3}
	//fmt.Println(l[4])

	// nil 不能当key
	//m[nil] = "a"
	//m["nil"] = nil

	testInterfaceNil()

	//testMap2()

	v := ComputeIfAbsent(m, "test", c)

	fmt.Println(v)
}

func ComputeIfAbsent(m map[string]string, k string, f func(key string) string) string {
	v, ok := m[k]
	if ok {
		return v
	}
	ans := f(k)
	m[k] = ans
	return ans
}

func testInterfaceNil() {
	var p *int = nil
	var i interface{} = p

	fmt.Printf("%v %+v is nil %v\n", i, asInterfaceStructure(i), i == nil)

}

// InterfaceStructure 定义了一个interface{}的内部结构
type InterfaceStructure struct {
	pt uintptr // 到值类型的指针
	pv uintptr // 到值内容的指针
}

func asInterfaceStructure(i interface{}) InterfaceStructure {
	return *(*InterfaceStructure)(unsafe.Pointer(&i))
}

func testMap2() {
	var m map[string]int

	//m["one"] = 1

	key := "two"
	elem, ok := m["two"]
	fmt.Printf("The element paired with key %q in nil map: %d (%v)\n",
		key, elem, ok)

	fmt.Printf("The length of nil map: %d\n", len(m))

	fmt.Printf("Delete the key-element pair by key %q...\n", key)
	delete(m, key)

	elem = 2
	fmt.Println("Add a key-element pair to a nil map...")
	m["two"] = elem // 这里会引发panic。
}
