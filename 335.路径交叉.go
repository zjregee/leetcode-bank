/*
 * @lc app=leetcode.cn id=335 lang=golang
 *
 * [335] 路径交叉
 */

// @lc code=start
func isSelfCrossing(distance []int) bool {
    for i := 3; i < len(distance); i++ {
        // 第 1 类路径交叉的情况
        if distance[i] >= distance[i-2] && distance[i-1] <= distance[i-3] {
            return true
        }

        // 第 2 类路径交叉的情况
        if i == 4 && distance[3] == distance[1] &&
            distance[4] >= distance[2]-distance[0] {
            return true
        }

        // 第 3 类路径交叉的情况
        if i >= 5 && distance[i-3]-distance[i-5] <= distance[i-1] &&
            distance[i-1] <= distance[i-3] &&
            distance[i] >= distance[i-2]-distance[i-4] &&
            distance[i-2] > distance[i-4] {
            return true
        }
    }
    return false
}
// @lc code=end

