/*
 * @lc app=leetcode.cn id=399 lang=golang
 *
 * [399] 除法求值
 *
 * https://leetcode-cn.com/problems/evaluate-division/description/
 *
 * algorithms
 * Medium (59.23%)
 * Likes:    535
 * Dislikes: 0
 * Total Accepted:    36.7K
 * Total Submissions: 62K
 * Testcase Example:  '[["a","b"],["b","c"]]\n' +
  '[2.0,3.0]\n' +
  '[["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]'
 *
 * 给你一个变量对数组 equations 和一个实数值数组 values 作为已知条件，其中 equations[i] = [Ai, Bi] 和
 * values[i] 共同表示等式 Ai / Bi = values[i] 。每个 Ai 或 Bi 是一个表示单个变量的字符串。
 *
 * 另有一些以数组 queries 表示的问题，其中 queries[j] = [Cj, Dj] 表示第 j 个问题，请你根据已知条件找出 Cj / Dj
 * = ? 的结果作为答案。
 *
 * 返回 所有问题的答案 。如果存在某个无法确定的答案，则用 -1.0 替代这个答案。如果问题中出现了给定的已知条件中没有出现的字符串，也需要用 -1.0
 * 替代这个答案。
 *
 * 注意：输入总是有效的。你可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries =
 * [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
 * 输出：[6.00000,0.50000,-1.00000,1.00000,-1.00000]
 * 解释：
 * 条件：a / b = 2.0, b / c = 3.0
 * 问题：a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
 * 结果：[6.0, 0.5, -1.0, 1.0, -1.0 ]
 *
 *
 * 示例 2：
 *
 *
 * 输入：equations = [["a","b"],["b","c"],["bc","cd"]], values = [1.5,2.5,5.0],
 * queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
 * 输出：[3.75000,0.40000,5.00000,0.20000]
 *
 *
 * 示例 3：
 *
 *
 * 输入：equations = [["a","b"]], values = [0.5], queries =
 * [["a","b"],["b","a"],["a","c"],["x","y"]]
 * 输出：[0.50000,2.00000,-1.00000,-1.00000]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * equations[i].length == 2
 * 1 i.length, Bi.length
 * values.length == equations.length
 * 0.0 < values[i]
 * 1
 * queries[i].length == 2
 * 1 j.length, Dj.length
 * Ai, Bi, Cj, Dj 由小写英文字母与数字组成
 *
 *
*/

// @lc code=start
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	dict := map[string]map[string]float64{}
	n := len(equations)
	for i := 0; i < n; i++ {
		res := float64(values[i])
		from, to := equations[i][0], equations[i][1]
		if _, ok := dict[from]; !ok {
			dict[from] = map[string]float64{from: 1.0}
		}
		dict[from][to] = res
		if _, ok := dict[to]; !ok {
			dict[to] = map[string]float64{to: 1.0}
		}
		dict[to][from] = 1 / res
	}
	var dfs func(visited map[string]bool, curr, root string)
	dfs = func(visited map[string]bool, curr, root string) {
		for k, v := range dict[curr] {
			if !visited[k] {
				visited[k] = true
				if _, ok := dict[root][k]; !ok {
					dict[root][k] = dict[root][curr] * v
				}
				dfs(visited, k, root)
			}
		}
	}
	for k := range dict {
		visited := map[string]bool{}
		dfs(visited, k, k)
	}
	m := len(queries)
	ans := make([]float64, m)
	for i, query := range queries {
		if v, ok := dict[query[0]]; ok {
			if u, ok1 := v[query[1]]; ok1 {
				ans[i] = u
				continue
			}
		}
		ans[i] = -1.0
	}
	return ans
}

// @lc code=end

