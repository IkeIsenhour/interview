package problem

import (
	"testing"

	utilityGoTest "github.com/IkeIsenhour/utility-go/pkg/test"
)

func TestContainsDuplicate(t *testing.T) {

	t.Run("no duplicates exist", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		got := ContainsDuplicate(nums)
		utilityGoTest.AssertEquality(t, got, false)
	})

	t.Run("duplicates exist", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 1}
		got := ContainsDuplicate(nums)
		utilityGoTest.AssertEquality(t, got, true)
	})
}
