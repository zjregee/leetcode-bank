/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m, n int) int {
    return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}
// @lc code=end

