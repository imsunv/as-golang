package main

import (
	"fmt"
	"testing"
	"unsafe"
)

func asInterfaceStructure(i interface{}) InterfaceStructure {
	return *(*InterfaceStructure)(unsafe.Pointer(&i))
}

// InterfaceStructure 定义了一个interface{}的内部结构
type InterfaceStructure struct {
	pt uintptr // 到值类型的指针
	pv uintptr // 到值内容的指针
}

func Test_InterfaceNil(t *testing.T) {
	var p *int = nil
	var i interface{} = p

	fmt.Printf("%v %+v is nil %v\n", i, asInterfaceStructure(i), i == nil)

	var p2 int = 2
	var i2 interface{} = p2

	fmt.Printf("%v %+v is nil %v\n", i2, asInterfaceStructure(i2), i2 == nil)
}
