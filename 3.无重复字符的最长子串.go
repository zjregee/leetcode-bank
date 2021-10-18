/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	n := len(s)
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for rk + 1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		ans = max(ans, rk - i + 1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

