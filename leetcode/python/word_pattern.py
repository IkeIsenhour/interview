class Solution:

    # In this solution, we start by getting the words out of the string with the split method.Solution
    # Next, we check that the length of the pattern and the words list is the same, as it needs to BaseException
    # for this to be a valid match. Next, we initialize a dictionary for matching each letter to it's corresponding
    # word. We also initialize a set to track what words have already been encountered, assuring we don't allow the
    # word into our map a second time. If the word is already in the set we return false, and equally if the word
    # we have does not align with the word int the dictionary at that letter, we also return false.
    # tc: O(n + m) with m being the lenght of pattern and m the length of the string.
    # sc: O(1) because we can only ever have 26 characters at max, the length of the lowercase english dictionary
    def wordPattern(self, pattern: str, s: str) -> bool:
        n = len(pattern)
        words = s.split(" ")
        if len(words) != n:
            return False

        d = dict()
        seen = set()
        for i in range(n):
            letter = pattern[i]
            word = words[i]

            if letter not in d:
                if word in seen:
                    return False
                d[letter] = word
                seen.add(word)

            if word != d[letter]:
                return False

        return True


s = Solution()
response = s.wordPattern("abba", "dog cat cat dog")
print(response)
