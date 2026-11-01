package problem

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	ans := make([]int, 2)

	for key, num := range nums {
		compliment := target - num

		if value, ok := m[compliment]; ok {
			ans[0] = key
			ans[1] = value
			break
		}

		m[num] = key
	}
	return ans
}
