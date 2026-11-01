package problem

import (
	"testing"

	utilityGoTest "github.com/IkeIsenhour/utility-go/pkg/test"
)

func TestClimbStairs(t *testing.T) {
	t.Run("success #1", func(t *testing.T) {
		got := ClimbStairs(2)
		want := 2
		utilityGoTest.AssertEquality(t, got, want)
	})
}
