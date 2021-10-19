/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func swapColors(colors []int, target int) (countTarget int) {
    for i, c := range colors {
        if c == target {
            colors[i], colors[countTarget] = colors[countTarget], colors[i]
            countTarget++
        }
    }
    return
}

func sortColors(nums []int) {
    count0 := swapColors(nums, 0) // 把 0 排到前面
    swapColors(nums[count0:], 1)  // nums[:count0] 全部是 0 了，对剩下的 nums[count0:] 把 1 排到前面
}
// @lc code=end

