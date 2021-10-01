/*
 * @lc app=leetcode.cn id=1720 lang=golang
 *
 * [1720] 解码异或后的数组
 */

// @lc code=start
func decode(encoded []int, first int) []int {
	ans := make([]int, len(encoded) + 1)
	ans[0] = first
	for i, e := range encoded {
		ans[i+1] = ans[i] ^ e
	}
	return ans
}
// @lc code=end

