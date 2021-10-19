/*
 * @lc app=leetcode.cn id=581 lang=golang
 *
 * [581] 最短无序连续子数组
 */

// @lc code=start
func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	minn, maxn := math.MaxInt64, math.MinInt64
	left, right := -1, -1
	for i, num := range nums {
		if maxn > num {
			right = i
		} else {
			maxn = num
		}
		if minn < nums[n-i-1] {
			left = n - i - 1
		} else {
			minn = nums[n-i-1]
		}
	}
	if right == -1 {
		return 0
	}
	return right - left + 1
}

// @lc code=end

