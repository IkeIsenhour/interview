package problem

import (
	"testing"

	utilityGoTest "github.com/IkeIsenhour/utility-go/pkg/test"
)

func TestProductOfArrayexceptSelf(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		got := ProductOfArrayExceptSelf(nums)
		want := []int{120, 60, 40, 30, 24}

		utilityGoTest.AssertSliceEquality(t, got, want)
	})

	t.Run("success with 0", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 0}
		got := ProductOfArrayExceptSelf(nums)
		want := []int{0, 0, 0, 0, 24}

		utilityGoTest.AssertSliceEquality(t, got, want)
	})
}
