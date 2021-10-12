/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	// var dfs func(*TreeNode, int, []int)
	dfs := func (root *TreeNode, targetSum int, path []int) {
		if root == nil {
			return
		}
		paths := make([]int, len(path))
		copy(paths, path)
		paths = append(paths, root.Val)
		if root.Left == nil && root.Right == nil && root.Val == targetSum {
			// copypath := make([]int, len(path))
			// copy(copypath, path)
			// res = append(res, copypath)
			res = append(res, paths)
		} else {
			dfs(root.Right, targetSum - root.Val, paths)
			dfs(root.Left, targetSum - root.Val, paths)
		}
	}
	dfs(root, targetSum, []int{})
	return res
}
// @lc code=end

