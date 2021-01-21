package assert

import (
	"testing"
)

func TestAssertEquals(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		Equals(t, "", 1, 2)
	})

	t.Run("string", func(t *testing.T) {
		Equals(t, "string: ", "aaa", "bbb")
	})
}
