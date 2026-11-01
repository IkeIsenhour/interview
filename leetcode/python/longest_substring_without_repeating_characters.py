def lengthOfLongestSubstring(s: str) -> int:
    longest = 0
    left = 0
    sett = set()

    for right in range(len(s)):
        while s[right] in sett:
            sett.remove(s[left])
            left += 1

        w = (right - left) + 1
        longest = max(longest, w)
        sett.add(s[right])

    return longest


print(lengthOfLongestSubstring("abab"))
print(lengthOfLongestSubstring("abcdbe"))
