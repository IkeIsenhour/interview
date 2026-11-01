package problem

func MaxProductSubarray(nums []int) int {
	res := nums[0]
	curMin, curMax := 1, 1

	for _, n := range nums {
		copy := curMax
		curMax = max(curMax*n, curMin*n, n)
		curMin = min(copy*n, curMin*n, n)

		res = max(curMax, res)
	}

	return res
}
