package problem

func ClimbStairs(n int) int {
	memo := map[int]int{
		1: 1,
		2: 2,
	}

	var f func(n int) int
	f = func(n int) int {
		val, ok := memo[n]
		if ok {
			return val
		} else {
			memo[n] = f(n-1) + f(n-2)
		}
		return memo[n]
	}

	return f(n)
}
