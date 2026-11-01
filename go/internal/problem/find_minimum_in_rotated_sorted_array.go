package problem

func FindMinimumInRotatedSortedArray(nums []int) int {
	res := nums[0]
	l, r := 0, len(nums)-1

	for l <= r {
		// fmt.Printf("\nl: %v, r: %v", l, r)
		if nums[l] < nums[r] {
			res = min(res, nums[l])
			break
		}

		m := (l + r) / 2
		res = min(res, nums[m])
		// fmt.Printf(", m: %v, res: %v", m, res)

		if nums[m] >= nums[l] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return res
}
