class Solution:
    def isHappy(self, n: int) -> bool:
        # setup set for tracking seen numbers
        seen = set()
        # convert the current number to a string so we can loop through it's digits
        curr = str(n)

        # continue looping until we've seen a number twice
        while curr not in seen:
            # add curr string
            seen.add(curr)
            # setup sum for tracking the total of the current digits
            summ = 0

            # loop through each digit in curr string
            for digit in curr:
                # convert char to int
                digit = int(digit)
                # square that digit and add it to sum
                summ += digit**2

            # check if sum is equal to 1, return true if so
            if summ == 1:
                return True
            # set cur value to new sum number
            curr = str(summ)

        # if we've made it here, that means we broke out of the loop via the seen set so return false
        return False
