package problem

import (
	"fmt"

	"github.com/IkeIsenhour/leetcode-go/internal/algorithm/search"
)

func SearchInRotatedSortedArray(nums []int, target int) int {
	pivotIdx := 0
	l, r := 0, len(nums)-1

	if nums[l] < nums[r] {
		return 0
	}

	for l <= r {
		if nums[l] < nums[r] {
			pivotIdx = l
			break
		}

		m := (l + r) / 2
		if nums[m] < nums[pivotIdx] {
			pivotIdx = m
		}

		if nums[m] >= nums[l] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	if target > nums[pivotIdx] && target > nums[len(nums)-1] {
		l = 0
		r = pivotIdx
	} else {
		l = pivotIdx
		r = len(nums) - 1
	}

	fmt.Printf("\n\np: %v, l: %v, r: %v", pivotIdx, l, r)

	return search.BinarySearch(nums[l:r], target)
}
