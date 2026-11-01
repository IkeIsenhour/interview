package problem

import (
	utilityGoTest "github.com/IkeIsenhour/utility-go/pkg/test"
	"testing"
)

func TestPow(t *testing.T) {
	t.Run("success positive", func(t *testing.T) {
		x := 2.0
		n := 10

		got := MyPow(x, n)
		want := 1024.0

		utilityGoTest.AssertEquality(t, got, want)
	})

	t.Run("success negative", func(t *testing.T) {
		x := 2.0
		n := -3

		got := MyPow(x, n)
		want := 1.0 / 8.0

		utilityGoTest.AssertEquality(t, got, want)
	})

	t.Run("success positive", func(t *testing.T) {
		x := 2.0
		n := 0

		got := MyPow(x, n)
		want := 1.0

		utilityGoTest.AssertEquality(t, got, want)
	})
}
