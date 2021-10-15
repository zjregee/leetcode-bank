/*
 * @lc app=leetcode.cn id=1690 lang=golang
 *
 * [1690] 石子游戏 VII
 */

// @lc code=start
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func stoneGameVII(stones []int) int {
    sums := make([]int, len(stones)+1)
    for i := 1; i <= len(stones); i++ {
        sums[i] = sums[i-1] + stones[i-1]
    }

    dp := make([][]int, len(stones))
    for i := 0; i < len(stones); i++ {
        dp[i] = make([]int, len(stones))
    }

    for i := len(stones) - 2; i >= 0; i-- {
        for j := i+1; j < len(stones); j++ {
            dp[i][j] = max(sums[j+1] - sums[i+1] - dp[i+1][j], sums[j] - sums[i] - dp[i][j-1])
        }
    }
    return dp[0][len(stones) - 1]
}
// @lc code=end

