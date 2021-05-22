/*
 * @lc app=leetcode.cn id=372 lang=golang
 *
 * [372] 超级次方
 *
 * https://leetcode-cn.com/problems/super-pow/description/
 *
 * algorithms
 * Medium (49.11%)
 * Likes:    114
 * Dislikes: 0
 * Total Accepted:    11.6K
 * Total Submissions: 23.4K
 * Testcase Example:  '2\n[3]'
 *
 * 你的任务是计算 a^b 对 1337 取模，a 是一个正整数，b 是一个非常大的正整数且会以数组形式给出。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：a = 2, b = [3]
 * 输出：8
 *
 *
 * 示例 2：
 *
 *
 * 输入：a = 2, b = [1,0]
 * 输出：1024
 *
 *
 * 示例 3：
 *
 *
 * 输入：a = 1, b = [4,3,3,8,5,2]
 * 输出：1
 *
 *
 * 示例 4：
 *
 *
 * 输入：a = 2147483647, b = [2,0,0]
 * 输出：1198
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 * 0
 * b 不含前导 0
 *
 *
 */

// @lc code=start
var base int = 1337

func pow(a, k int) int {
	if k == 0 {
		return 1
	}
	a %= base
	if k%2 == 0 {
		ans := pow(a, k/2)
		return (ans * ans) % base
	} else {
		ans := pow(a, (k-1)/2)
		return (a * (ans * ans) % base) % base
	}
}
func superPow(a int, b []int) int {
	n := len(b)
	if n == 0 {
		return 1
	}
	dp := 1
	for i := range b {
		dp = (pow(a, b[i]) * pow(dp, 10)) % base
	}
	return dp
}

// @lc code=end

