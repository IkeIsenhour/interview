package problem

import (
	"testing"

	utilityGoTest "github.com/IkeIsenhour/utility-go/pkg/test"
)

func TestFindMinimumInRotatedSortedArray(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		nums := []int{30, 41, 49, 56, 9, 19, 25}

		got := FindMinimumInRotatedSortedArray(nums)
		want := 9

		utilityGoTest.AssertEquality(t, got, want)
	})
}
