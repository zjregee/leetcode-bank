/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
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
// func rob(root *TreeNode) int {
// 	var dfs func (*TreeNode, int) int
// 	dfs = func(root *TreeNode, flag int) int {
// 		if root == nil {
// 			return 0
// 		}
// 		reward := 0
// 		if flag == 1 {
// 			reward = root.Val + dfs(root.Left, 0) + dfs(root.Right, 0)
// 		}
// 		reward = max(reward, dfs(root.Left, 1) + dfs(root.Right, 1))
// 		return reward
// 	}
// 	return dfs(root, 1)
// }

// func rob(root *TreeNode) int {
// 	choose := map[*TreeNode]int{}
// 	unchoose := map[*TreeNode]int{}

// 	var dfs func (*TreeNode, int) int
// 	dfs = func(root *TreeNode, flag int) int {
// 		if root == nil {
// 			return 0
// 		}
// 		reward1 := 0
// 		reward2 := 0
// 		if flag == 1 {
// 			if unchoose[root.Left] == 0 {
// 				unchoose[root.Left] = dfs(root.Left, 0)
// 			}
// 			if unchoose[root.Right] == 0 {
// 				unchoose[root.Right] = dfs(root.Right, 0)
// 			}
// 			reward1 = root.Val + unchoose[root.Left] + unchoose[root.Right]
// 			choose[root] = reward1
// 		}
// 		if choose[root.Left] == 0 {
// 			choose[root.Left] = dfs(root.Left, 1)
// 		}
// 		if choose[root.Right] == 0{
// 			choose[root.Right] = dfs(root.Right, 1)
// 		}
// 		reward2 = choose[root.Left] + choose[root.Right]
// 		unchoose[root] = reward2
// 		return max(reward1, reward2)
// 	}
// 	return dfs(root, 1)
// 	// return max(unchoose[root], choose[root])
// }

func rob(root *TreeNode) int {
	choose := map[*TreeNode]int{}
	unchoose := map[*TreeNode]int{}
	choose[nil] = 0
	unchoose[nil] = 0
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		dfs(root.Right)
		choose[root] = root.Val + unchoose[root.Left] + unchoose[root.Right]
		unchoose[root] = max(choose[root.Left], unchoose[root.Left]) + max(choose[root.Right], unchoose[root.Right])
	}
	dfs(root)
	return max(choose[root], unchoose[root])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

