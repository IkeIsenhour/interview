package problem

func MaxProfit(prices []int) int {
	fast := 1
	slow := 0
	maxDiff := 0

	for fast < len(prices) {
		if prices[slow] < prices[fast] {
			diff := prices[fast] - prices[slow]
			if diff > maxDiff {
				maxDiff = diff
			}
		} else {
			slow = fast
		}
		fast++
	}

	return maxDiff
}
