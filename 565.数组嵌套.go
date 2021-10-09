/*
 * @lc app=leetcode.cn id=565 lang=golang
 *
 * [565] 数组嵌套
 */

// @lc code=start
func arrayNesting(nums []int) int {
	visited := make([]bool, len(nums))
	res := 0
	for i := 0; i < len(nums); i++ {
		if !visited[i] {
			start := nums[i]
			count := 0
			for {
				start = nums[start]
				count++
				visited[start] = true
				if start == nums[i] {
					break
				}
			}
			res = max(res, count)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

