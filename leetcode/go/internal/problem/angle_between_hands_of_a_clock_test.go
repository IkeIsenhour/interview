package problem

import (
	"testing"

	utilityGoTest "github.com/IkeIsenhour/utility-go/pkg/test"
)

func TestAngleClock(t *testing.T) {

	t.Run("check the angle of 3:15", func(t *testing.T) {
		got := AngleClock(3, 15)
		want := 7.50

		utilityGoTest.AssertEquality(t, got, want)
	})

	t.Run("check the angle when hands are in the same location", func(t *testing.T) {
		got := AngleClock(12, 0)
		want := 0.0

		utilityGoTest.AssertEquality(t, got, want)
	})
}
