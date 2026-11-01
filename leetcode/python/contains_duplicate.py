def contains_duplicate(nums: list[int]) -> bool:
    if len(nums) == 0:
        return False

    seen = set()
    for num in nums:
        if num in seen:
            return True
        seen.add(num)

    return False


test_one = [1, 2, 3, 4, 5, 6]
test_two = [1, 2, 3, 4, 5, 1]

print(contains_duplicate(test_one))
print(contains_duplicate(test_two))
