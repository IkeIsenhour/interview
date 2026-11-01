package search

import "fmt"

func BinarySearch(nums []int, target int) int {

	fmt.Println(nums)
	low, high := 0, len(nums)-1
	var mid int

	for low <= high {
		mid = low + (high-low)/2

		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
