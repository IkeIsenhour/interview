def missing_number(nums: list[int]) -> int:
    number = len(nums)
    total = (number * (number + 1)) // 2

    for num in nums:
        total -= num

    return total


test_one = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 11]  ## Missing 10
test_two = [0, 3, 2, 5, 1]  ## Missing 4

print(missing_number(test_one))
print(missing_number(test_two))
