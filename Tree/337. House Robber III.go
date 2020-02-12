/*
337. House Robber III
The thief has found himself a new place for his thievery again. There is only one entrance to this area, called the "root." Besides the root, each house has one and only one parent house. After a tour, the smart thief realized that "all houses in this place forms a binary tree". It will automatically contact the police if two directly-linked houses were broken into on the same night.

Determine the maximum amount of money the thief can rob tonight without alerting the police.

Example 1:

Input: [3,2,3,null,3,null,1]

     3
    / \
   2   3
    \   \ 
     3   1

Output: 7 
Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.
Example 2:

Input: [3,4,5,1,3,null,1]

     3
    / \
   4   5
  / \   \ 
 1   3   1

Output: 9
Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.
*/
/*
Solution:
for every node, there are 2 actions: rob or not.  return values reprensent the actions. 0: not rob. 1: rob.
if robbing root, sum is rob root, not rob left and not rob right.
*/

package main
import (
	"fmt"
)
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func rob(root *TreeNode) int {
	if nil == root {
		return 0
	}
	res := dfs(root)
	return max(res[0], res[1])
}

func dfs(root *TreeNode) []int {
	if nil == root {
		return []int{0, 0}
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	rob := root.Val + left[0] + right[0]
	not_rob := max(left[0], left[1]) + max(right[0], right[1])
	return []int{not_rob, rob}
}