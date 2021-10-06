/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 */

// @lc code=start
func thirdMax(nums []int) int {
	var a, b, c *int
	for _, num := range nums {
		num := num
		if a == nil || num > *a {
			a, b, c = &num, a, b
		} else if *a > num && (b == nil || num > *b) {
			b, c = &num, b
		} else if b != nil && *b > num && (c == nil || num > *c) {
			c = &num
		}
	}
	if c == nil {
		return *a
	}
	return *c
}
// @lc code=end

