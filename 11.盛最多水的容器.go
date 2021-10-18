/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	ans := 0
	for {
		area := min(height[l], height[r]) * (r - l)
		ans = max(ans, area)
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
		if l >= r {
			break
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

