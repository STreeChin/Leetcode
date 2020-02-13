/*
76. Minimum Window Substring
Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).
Example: Input: S = "ADOBECODEBANC", T = "ABC"  Output: "BANC"
*/
/*solution
（1）sliding window. Map the byte in s in advance.
（2）The default value of the map is 0, use it as the flag.
*/
func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) > len(s) {
		return ""
	}
	if len(t) == 0 {
		return s[:1]
	}
	left, right, start, match, minLen := 0, 0, 0, 0, len(s)+1
	tMap := map[byte]int{}
	for _, v := range t {
		tMap[byte(v)]++
	}
	for ; right < len(s); right++ {
		if tMap[s[right]]--; tMap[s[right]] >= 0 {
			match++
		}

		for ; match == len(t); left++ { 
			if right-left < minLen {
				start = left
				minLen = right - left
			}
			if tMap[s[left]]++; tMap[s[left]] > 0 {
				match--
			}
		}
	}
	if minLen+start+1 > len(s) {
		return ""
	} else {
		return s[start : minLen+start+1]
	}
}