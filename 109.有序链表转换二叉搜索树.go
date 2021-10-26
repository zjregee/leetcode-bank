/*
 * @lc app=leetcode.cn id=109 lang=golang
 *
 * [109] 有序链表转换二叉搜索树
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var globalHead *ListNode

func sortedListToBST(head *ListNode) *TreeNode {
	globalHead = head
	length := getLength(head)
	return buildTree(0, length-1)
}

func getLength(head *ListNode) int {
	ret := 0
	for ; head != nil; head = head.Next {
		ret++
	}
	return ret
}

func buildTree(left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right + 1) / 2
	root := &TreeNode{}
	root.Left = buildTree(left, mid-1)
	root.Val = globalHead.Val
	globalHead = globalHead.Next
	root.Right = buildTree(mid+1, right)
	return root
}

// @lc code=end

