package problem

import (
	"testing"

	utilityGoTest "github.com/IkeIsenhour/utility-go/pkg/test"
)

func TestMaxProfit(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		stocks := []int{7, 1, 5, 3, 6, 4}
		got := MaxProfit(stocks)
		utilityGoTest.AssertEquality(t, got, 5)
	})
}
