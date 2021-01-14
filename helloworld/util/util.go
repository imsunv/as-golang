package util2

import "fmt"

func Echo(input string) string {
	if IsRight() {
		fmt.Println("Right")
	}
	return "Echo: " + input + " "
}
