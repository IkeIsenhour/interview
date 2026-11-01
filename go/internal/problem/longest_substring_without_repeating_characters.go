package problem

// LengthOfLongestSubstringNestedLoopHashMap iterate over each
// byte in the string. One loop for each character, then a nested loop for each character
// after that character. Once in the top level loop, instantiate the a boolean hash map
// to track if the character has been found in this iteration. This
// would probably be considered the brute force solution.
// tc: O(n^2), sc: O(n)
func LengthOfLongestSubstringNestedLoopHashMap(s string) int {
	if len(s) == 0 {
		return 0
	}

	longestLen := 1
	for i := 0; i < len(s); i++ {
		charMap := map[byte]bool{s[i]: true}
		currLen := 1
		for j := i + 1; j < len(s); j++ {
			currChar := s[j]
			_, exists := charMap[currChar]
			if exists {
				break
			}
			charMap[currChar] = true
			currLen++
		}

		if currLen > longestLen {
			longestLen = currLen
		}
	}

	return longestLen
}

// LengthOfLongestSubstringSlidingWindow uses the sliding window technique and a hashmap to come to
// the optimize solution. For an in-depth explanation, reference the problemSolutionMethods.md in
// the docs directory.
// tc: O(n), sc: O(n)
func LengthOfLongestSubstringSlidingWindow(s string) int {
	left, right, longest := 0, 0, 0
	set := map[byte]bool{}

	for right < len(s) {
		if _, exists := set[s[right]]; exists {
			delete(set, s[left])
			left++
		} else {
			set[s[right]] = true
			if w := (right - left) + 1; w > longest {
				longest = w
			}
			right++
		}
	}

	return longest
}
