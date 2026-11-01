package problem

func MaxSubArray(nums []int) int {
	ans := nums[0]
	currSum := 0

	for _, num := range nums {
		if currSum < 0 {
			currSum = 0
		}

		currSum += num
		ans = max(currSum, ans)
	}

	return ans
}
