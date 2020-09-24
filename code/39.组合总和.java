import java.util.ArrayList;
import java.util.List;

/*
 * @lc app=leetcode.cn id=39 lang=java
 *
 * [39] 组合总和
 *
 * https://leetcode-cn.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (66.95%)
 * Likes:    837
 * Dislikes: 0
 * Total Accepted:    126.4K
 * Total Submissions: 181.8K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 * 
 * candidates 中的数字可以无限制重复被选取。
 * 
 * 说明：
 * 
 * 
 * 所有数字（包括 target）都是正整数。
 * 解集不能包含重复的组合。 
 * 
 * 
 * 示例 1：
 * 
 * 输入：candidates = [2,3,6,7], target = 7,
 * 所求解集为：
 * [
 * ⁠ [7],
 * ⁠ [2,2,3]
 * ]
 * 
 * 
 * 示例 2：
 * 
 * 输入：candidates = [2,3,5], target = 8,
 * 所求解集为：
 * [
 * [2,2,2,2],
 * [2,3,3],
 * [3,5]
 * ]
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= candidates.length <= 30
 * 1 <= candidates[i] <= 200
 * candidate 中的每个元素都是独一无二的。
 * 1 <= target <= 500
 * 
 * 
 */

// @lc code=start
class Solution {
    
    private int dfs(List<List<Integer>> ans, int[] candidates, int target, int index, List<Integer> combination) {
        if (target == 0) {
            ans.add(new ArrayList<Integer>(combination));
            return 1;
        }

        if (target < 0) {
            return -1;
        }

        for(int i = index; i < candidates.length; i++) {
            combination.add(candidates[i]);
            int res = dfs(ans, candidates, target - candidates[i], i, combination);
            combination.remove(combination.size() - 1);
            // terminate early if target < 0
            if (res == -1) {
                return 0;
            }
        }
        return 0;
    }
    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        List<List<Integer>> ans = new ArrayList<>();
        if (candidates.length == 0) return ans;
        Arrays.sort(candidates);
        
        dfs(ans, candidates, target, 0, new ArrayList<Integer>());

        return ans;
    }
}
// @lc code=end

