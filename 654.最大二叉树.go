/*
 * @lc app=leetcode.cn id=654 lang=golang
 *
 * [654] 最大二叉树
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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	// pos := 0
	// max := 0
	// for i, num := range nums {
	// 	if num > pos {
	// 		max = num
	// 		pos = i
	// 	}
	// }
	// node := &TreeNode{}
	// node.Val = max
	// if pos != 0 {
	// 	node.Left = constructMaximumBinaryTree(nums[0:pos-1])
	// }
	// if pos + 1 != len(nums) {
	// 	node.Right = constructMaximumBinaryTree(nums[pos:])
	// }
	// return node
	return construct(nums, 0, len(nums))
}

func construct(nums []int, l, r int) *TreeNode {
	if l == r {
		return nil
	}
	pos := 0
	for i := l; i < r; i++ {
		if nums[i] > nums[pos] {
			pos = i
		}
	}
	node := &TreeNode{}
	node.Val = nums[pos]
	node.Left = construct(nums, l, pos)
	node.Right = construct(nums, pos + 1, r)
	return node
}
// @lc code=end

