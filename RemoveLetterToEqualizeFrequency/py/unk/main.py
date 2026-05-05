# https://leetcode.com/problems/remove-letter-to-equalize-frequency/description/


class Solution(object):
    def equalFrequency(self, word: str):
        """
        :type word: str
        :rtype: bool
        """
        length = len(word)
        if length <= 1:
            return True

        for i, c in enumerate(word):
            newWord = word[:i] + word[i + 1 :]

            m = {}

            for j, c2 in enumerate(newWord):
                m[c2] = m.get(c2, 0) + 1

            values = m.values()
            prev = list(values)[0]
            failed = False
            for value in m.values():
                if prev != value:
                    failed = True
                prev = value
            if failed:
                continue
            else:
                return True
        return False


if __name__ == "__main__":
    solution = Solution()

    # Example 1: "abcc"
    # Remove 'c' -> "abc", all characters have frequency 1
    assert solution.equalFrequency("abcc") == True

    # Example 2: "aab"
    # Remove 'a' -> "ab", both have frequency 1
    assert solution.equalFrequency("aab") == True

    # Example 3: "aaabbbcc"
    # No valid removal
    assert solution.equalFrequency("aaabbbcc") == False

    # Example 4: "a"
    # Single character, already equal frequency
    assert solution.equalFrequency("a") == True

    # Example 5: "ab"
    # Remove either character, remaining has frequency 1
    assert solution.equalFrequency("ab") == True

    # Example 6: "abc"
    # All have frequency 1, remove any makes others unequal or valid
    assert solution.equalFrequency("abc") == True

    # Example 7: "aaaa"
    # Remove one 'a' -> "aaa", all same frequency
    assert solution.equalFrequency("aaaa") == True

    # Example 8: "aabbcc"
    # All frequencies are 2, removing one makes it unequal
    assert solution.equalFrequency("aabbcc") == False

    print("All test cases passed!")
