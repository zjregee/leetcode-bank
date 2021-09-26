/*
 * @lc app=leetcode.cn id=1928 lang=golang
 *
 * [1928] 规定时间内到达终点的最小花费
 */

// @lc code=start
func minCost(maxTime int, edges [][]int, passingFees []int) int {
	inf := 10000*101 + 1  // 极大值
	n := len(passingFees) // 城市数量
	f := make([][]int, maxTime+1)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = inf
		}
	}
	f[0][0] = passingFees[0]
	for t := 1; t <= maxTime; t++ {
		for _, edge := range edges {
			j, i, time := edge[0], edge[1], edge[2]
			if time > t {
				continue
			}
			f[t][i] = min(f[t][i], f[t-time][j] + passingFees[i])
			f[t][j] = min(f[t][j], f[t-time][i] + passingFees[j])
		}
	}
	ans := inf
	for t := 1; t <= maxTime; t++ {
		ans = min(ans, f[t][n-1])
	}
	if ans == inf {
		ans = -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

