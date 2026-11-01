class Solution:
    # This brute force solution add each number to a dictionary where it trakcs
    # the occurances of that number.
    # tc: O(n) because it loops through the list a single time
    # sc: O(n) because it stores each number 1 time in a map
    def majorityElementBruteForce(self, nums: List[int]) -> int:
        counter = dict()

        for num in nums:
            if num in counter:
                counter[num] += 1
            else:
                counter[num] = 1

            if counter[num] >= (len(nums) / 2):
                return num

    # This is the optimal solution, uses an algorithm that essentially takes a running count that is inremented/decremented
    # based on what the current answer is.
    # Explanation: Majority Element - Leetcode 169 - Hashmaps & Sets (Python) , Greg Hogg https://www.youtube.com/watch?v=c1B3LZQtZ_s
    # tc: O(n)
    # sc: O(1)
    def majorityElement(self, nums: List[int]) -> int:
        ans = -1
        count = 0

        for num in nums:
            if count == 0:
                ans = num
            if num == ans:
                count += 1
            else:
                count -= 1

        return ans
