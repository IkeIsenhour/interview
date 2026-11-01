# Nested for loop
class Solution:
    def twoSumBrute(self, nums: List[int], target: int) -> List[int]:
        for i in range(len(nums)):
            for j in range(i + 1, len(nums)):
                if nums[i] + nums[j] == target:
                    return [i, j]


# Find the complement and check for it in a hash map
class Solution:
    def twoSumOptimized(self, nums: List[int], target: int) -> List[int]:
        map = {}
        for i in range(len(nums)):
            curr = nums[i]
            comp = target - curr
            if comp in map:
                return [i, map[comp]]
            else:
                map[curr] = i
        return []
