package problem

func ContainsDuplicate(nums []int) bool {
	m := make(map[int]bool)

	for _, v := range nums {
		if _, exists := m[v]; exists {
			return true
		}
		m[v] = true
	}

	return false
}
