package main

import (
	"fmt"
	util "github.com/imsunv/as-golang/helloworld/util"
)

func main() {
	if util.IsRight() {
		fmt.Println(MustSamePackageInSameDirectory())
	}

	fmt.Println(util.Echo("Hello World"))
}
