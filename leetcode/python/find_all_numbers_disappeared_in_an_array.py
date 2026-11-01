def find_all_missing_numbers_brute(nums: list[int]) -> list[int]:
    nums_set = set(nums)
    missing = []

    for i in range(1, len(nums) + 1):
        if i not in nums_set:
            missing.append(i)
    return missing


def find_all_missing_numbers_optimized_no_extra_memory(nums: list[int]) -> list[int]:
    for n in nums:
        i = abs(n) - 1
        nums[i] = -1 * abs(nums[i])

    missing = []
    for i, n in enumerate(nums):
        if n > 0:
            missing.append(i + 1)

    return missing
