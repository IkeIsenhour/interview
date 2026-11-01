package problem

import (
	"testing"

	utilityGoTest "github.com/IkeIsenhour/utility-go/pkg/test"
)

func TestTwoSum(t *testing.T) {
	arr := []int{2, 21, 303, 7}

	t.Run("successful TwoSum #1", func(t *testing.T) {
		got := TwoSum(arr, 9)
		want := []int{3, 0}

		utilityGoTest.AssertSliceEquality(t, got, want)
	})
}
