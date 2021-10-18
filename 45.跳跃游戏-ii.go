/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	n := len(nums)
	rightmost := 0
	for i := 0; i < n; i++ {
		rightmost = max(rightmost, i + nums[i]);
		if rightmost >= n - 1 {
			return true
		}
	}
	return false
}
// @lc code=end

