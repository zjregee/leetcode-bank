/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start
func fizzBuzz(n int) (ans []string) {
	for i := 1; i <= n; i++ {
		sb := &strings.Builder{}
		if i % 3 == 0 {
			sb.WriteString("Fizz")
		}
		if i % 5 == 0 {
			sb.WriteString("Buzz")
		}
		if sb.Len() == 0 {
			sb.WriteString(strconv.Itoa(i))
		}
		ans = append(ans, sb.String())
	}
	return
}
// @lc code=end

