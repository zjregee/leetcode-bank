/*
 * @lc app=leetcode.cn id=453 lang=golang
 *
 * [453] 最小操作次数使数组元素相等
 */

// @lc code=start
func minMoves(nums []int) (ans int) {
	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	for _, num := range nums {
		ans += num - min
	}
	return
}
// @lc code=end

