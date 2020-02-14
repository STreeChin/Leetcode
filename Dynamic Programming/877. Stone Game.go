/*
877. Stone Game
Alex and Lee play a game with piles of stones.  There are an even number of piles arranged in a row, and each pile has a positive integer number of stones piles[i].

The objective of the game is to end with the most stones.  The total number of stones is odd, so there are no ties.

Alex and Lee take turns, with Alex starting first.  Each turn, a player takes the entire pile of stones from either the beginning or the end of the row.  This continues until there are no more piles left, at which point the person with the most stones wins.

Assuming Alex and Lee play optimally, return True if and only if Alex wins the game.
*/
/*（1）solution

*/

func stoneGame(piles []int) bool {
	dp := []int{}
	dp = append(dp, piles...)
	for d := 1; d < len(piles); d++ {
		for i := 0; i < len(piles)-d; i++ {
			dp[i] = max(piles[i] - dp[i + 1], piles[i + d] - dp[i])
		}
	}
	return dp[0] > 0
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}