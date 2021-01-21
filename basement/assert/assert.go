package assert

import (
	"fmt"
	"testing"
)

func Equals(t testing.TB, message string, expected interface{}, actual interface{}) {
	if expected != actual {
		failNotEquals(t, message, expected, actual)
	}
}

func IsTrue(t testing.TB, actual bool) {
	Equals(t, "", true, actual)
}

func EqualsInt(t testing.TB, message string, expected int, actual int) {
	if expected != actual {
		failNotEquals(t, message, expected, actual)
	}
}

func failNotEquals(t testing.TB, message string, expected interface{}, actual interface{}) {
	t.Helper()
	t.Error(format(message, expected, actual))
}

func format(message string, expected interface{}, actual interface{}) string {
	formatted := ""
	if message != "" {
		formatted = message + " "
	}

	return fmt.Sprintf(formatted+"expected:<%v> but was:<%v>", expected, actual)
}
