/*
 * @lc app=leetcode.cn id=310 lang=golang
 *
 * [310] 最小高度树
 *
 * https://leetcode-cn.com/problems/minimum-height-trees/description/
 *
 * algorithms
 * Medium (36.99%)
 * Likes:    323
 * Dislikes: 0
 * Total Accepted:    16.6K
 * Total Submissions: 45K
 * Testcase Example:  '4\n[[1,0],[1,2],[1,3]]'
 *
 * 树是一个无向图，其中任何两个顶点只通过一条路径连接。 换句话说，一个任何没有简单环路的连通图都是一棵树。
 *
 * 给你一棵包含 n 个节点的树，标记为 0 到 n - 1 。给定数字 n 和一个有 n - 1 条无向边的 edges
 * 列表（每一个边都是一对标签），其中 edges[i] = [ai, bi] 表示树中节点 ai 和 bi 之间存在一条无向边。
 *
 * 可选择树中任何一个节点作为根。当选择节点 x 作为根节点时，设结果树的高度为 h 。在所有可能的树中，具有最小高度的树（即，min(h)）被称为
 * 最小高度树 。
 *
 * 请你找到所有的 最小高度树 并按 任意顺序 返回它们的根节点标签列表。
 * 树的 高度 是指根节点和叶子节点之间最长向下路径上边的数量。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 4, edges = [[1,0],[1,2],[1,3]]
 * 输出：[1]
 * 解释：如图所示，当根是标签为 1 的节点时，树的高度是 1 ，这是唯一的最小高度树。
 *
 * 示例 2：
 *
 *
 * 输入：n = 6, edges = [[3,0],[3,1],[3,2],[3,4],[5,4]]
 * 输出：[3,4]
 *
 *
 * 示例 3：
 *
 *
 * 输入：n = 1, edges = []
 * 输出：[0]
 *
 *
 * 示例 4：
 *
 *
 * 输入：n = 2, edges = [[0,1]]
 * 输出：[0,1]
 *
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * edges.length == n - 1
 * 0 i, bi < n
 * ai != bi
 * 所有 (ai, bi) 互不相同
 * 给定的输入 保证 是一棵树，并且 不会有重复的边
 *
 *
 */

// @lc code=start
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	} else if n == 2 {
		return []int{0, 1}
	}
	G := make([][]int, n)
	inDegree := make([]int, n)
	for _, edge := range edges {
		v := edge[0]
		u := edge[1]
		if G[v] == nil {
			G[v] = []int{u}
		} else {
			G[v] = append(G[v], u)
		}
		if G[u] == nil {
			G[u] = []int{v}
		} else {
			G[u] = append(G[u], v)
		}
		inDegree[u]++
		inDegree[v]++
	}
	queue := []int{}
	for i, v := range inDegree {
		if v == 1 {
			queue = append(queue, i)
		}
	}
	count := len(queue)
	for n > 2 {
		n -= count
		for ; count > 0; count-- {
			item := queue[0]
			queue = queue[1:]
			inDegree[item]--
			for _, v := range G[item] {
				inDegree[v]--
				if inDegree[v] == 1 {
					queue = append(queue, v)
				}
			}
		}
		count = len(queue)
	}
	ans := make([]int, 0, 2)
	for len(queue) > 0 {

		ans = append(ans, queue[len(queue)-1])
		queue = queue[:len(queue)-1]
	}

	return ans

}

// @lc code=end

