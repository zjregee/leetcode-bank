/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := rootSum(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}

func rootSum(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return
	}
	val := root.Val
	if val == targetSum {
		res++
	}
	res += rootSum(root.Left, targetSum-val)
	res += rootSum(root.Right, targetSum-val)
	return
}

// 前缀和
// func pathSum(root *TreeNode, targetSum int) (ans int) {
// 	preSum := map[int64]int{0: 1}
// 	var dfs func(*TreeNode, int64)
// 	dfs = func(node *TreeNode, curr int64) {
// 		if node == nil {
// 			return
// 		}
// 		curr += int64(node.Val)
// 		ans += preSum[curr-int64(targetSum)]
// 		preSum[curr]++
// 		dfs(node.Left, curr)
// 		dfs(node.Right, curr)
// 		preSum[curr]-- // 最重要的一行代码
// 		return
// 	}
// 	dfs(root, 0)
// 	return
// }
// @lc code=end

