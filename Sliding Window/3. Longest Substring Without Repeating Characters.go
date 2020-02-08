/*
Given a string, find the length of the longest substring without repeating characters.
Example 1: Input: "abcabcbb" Output: 3  Explanation: The answer is "abc", with the length of 3. 
Example 2: Input: "bbbbb"    Output: 1  Explanation: The answer is "b", with the length of 1.
Example 3: Input: "pwwkew"   Output: 3  Explanation: The answer is "wke", with the length of 3. 
          Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
/*solution
（1）The value of key in map is the maximum sequence number of the key in the string.
（2）set the start to -1, acting like the dummyhead of the link.
*/
func lengthOfLongestSubstring(s string) int {
	maxlenth := 0
	hmap := map[byte]int{}
	for i, start := 0, -1; i < len(s); i++ {
		if _, ok := hmap[s[i]]; ok {
			if hmap[s[i]] > start {
				start = hmap[s[i]]
			}
		}
		hmap[s[i]] = i
		if i-start > maxlenth {
			maxlenth = i - start
		}
	}
	return maxlenth
}