package function

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Addr() func(int) int {
	num := 0
	return func(i int) int {
		num += i
		return num
	}
}

type GenInt func() int

func fibonacci() GenInt {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 函数也是可以实现接口的
func (g GenInt) Read(p []byte) (n int, err error) {
	next := g()

	if next > 1000 {
		return 0, io.EOF
	}

	s := fmt.Sprintf("%d \n", next)
	return strings.NewReader(s).Read(p)
}

func PrintFibonacci(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		println(scanner.Text())
	}
}
