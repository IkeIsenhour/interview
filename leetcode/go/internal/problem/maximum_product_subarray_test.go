package problem

import (
	"testing"

	utilityGoTest "github.com/IkeIsenhour/utility-go/pkg/test"
)

func TestMaxProductSubarray(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		nums := []int{2, 3, -2, 4}

		got := MaxProductSubarray(nums)
		want := 6

		utilityGoTest.AssertEquality(t, got, want)
	})

	t.Run("Failure", func(t *testing.T) {
		nums := []int{2, 3, -2, 4}

		got := MaxProductSubarray(nums)
		want := 5

		utilityGoTest.AssertInequality(t, got, want)
	})
}
