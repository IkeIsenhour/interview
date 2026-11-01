package problem

func MyPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 1 / MyPow(x, -1*n)
	}

	if n%2 == 1 {
		return x * MyPow(x*x, (n-1)/2)
	}

	return MyPow(x*x, n/2)
}
