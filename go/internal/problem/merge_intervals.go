package problem

import (
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {
	// Sort slice
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Setup return slice
	output := [][]int{intervals[0]}

	// Loop through intervals
	for i := 1; i < len(intervals); i++ {
		lastEnd := output[len(output)-1][1]
		start := intervals[i][0]
		end := intervals[i][1]

		// Core logic
		if start <= lastEnd {
			output[len(output)-1][1] = int(math.Max(float64(lastEnd), float64(end)))
		} else {
			output = append(output, intervals[i])
		}
	}

	// Return
	return output
}
