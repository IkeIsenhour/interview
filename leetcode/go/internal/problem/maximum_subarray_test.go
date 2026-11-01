package problem

import (
	"testing"

	utilityGoTest "github.com/IkeIsenhour/utility-go/pkg/test"
)

func TestMaximumSubArray(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

		got := MaxSubArray(nums)
		want := 6

		utilityGoTest.AssertEquality(t, got, want)
	})

	t.Run("failure", func(t *testing.T) {
		nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

		got := MaxSubArray(nums)
		want := 1

		utilityGoTest.AssertInequality(t, got, want)
	})
}
