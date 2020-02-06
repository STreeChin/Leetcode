/*
516. Longest Palindromic Subsequence——medium
Given a string s, find the longest palindromic subsequence's length in s. You may assume that the maximum length of s is 1000.
Input:"bbbab"，Output:4。One possible longest palindromic subsequence is "bbbb".
Input:"cbbd"，Output:2。
*/
/*（1）solution
defination of dp ：For s[i..j], the length of Longest Palindromic Subsequence is dp[i][j]
（2）equation
if (s[i] == s[j])     // s[i],s[j] must be included in the Longest Palindromic Subsequence
    dp[i][j] = dp[i + 1][j - 1] + 2;
else
    dp[i][j] = max(dp[i + 1][j], dp[i][j - 1]);

dp[i][j] is calculated via dp[i+1][j-1]，dp[i+1][j]，dp[i][j-1].  
We can choose diagonal traversal or reverse traversal. The latter code is as follows
*/
package DP

import "fmt"

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func longestPalindromeSubseq(s string) int {
	if len(s) == 0 {
		return 0
	}
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][n-1]
}
