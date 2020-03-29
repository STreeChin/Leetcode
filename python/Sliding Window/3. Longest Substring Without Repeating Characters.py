
#Given a string, find the length of the longest substring without repeating characters.
#Example 1: Input: "abcabcbb" Output: 3  Explanation: The answer is "abc", with the length of 3. 
#Example 2: Input: "bbbbb"    Output: 1  Explanation: The answer is "b", with the length of 1.
#Example 3: Input: "pwwkew"   Output: 3  Explanation: The answer is "wke", with the length of 3. 
#          Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
#solution
#（1）The value of key in map is the maximum sequence number of the key in the string.
#（2）set the start to -1, acting like the dummyhead of the link.

def lengthOfLongestSubstring(s):
    maxlenth, start = 0, 0
    lookup = dict.fromkeys(s,0)
    for i in range(len(s)):
        start = max(start, lookup[s[i]])
        maxlenth = max(maxlenth, i - start)
        lookup[s[i]] = i
    return maxlenth