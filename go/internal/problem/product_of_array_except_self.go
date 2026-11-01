package problem

func ProductOfArrayExceptSelf(nums []int) []int {
	product := 1
	zeroFound := false

	for _, v := range nums {
		if v == 0 {
			if zeroFound {
				ans := make([]int, len(nums))
				return ans
			}
			zeroFound = true
		} else {
			product *= v
		}
	}

	ans := make([]int, len(nums))
	if zeroFound {
		for k, v := range nums {
			if v == 0 {
				ans[k] = product
			} else {
				ans[k] = 0
			}
		}
	} else {
		for k, v := range nums {
			ans[k] = product / v
		}
	}

	return ans
}
