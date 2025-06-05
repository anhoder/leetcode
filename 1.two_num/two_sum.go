// https://leetcode.cn/problems/two-sum/?envType=study-plan-v2&envId=top-100-liked
//
// 1. 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
//
// 你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
//
// 你可以按任意顺序返回答案。

package twosum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, v := range nums {
		left := target - v
		if j, ok := m[left]; ok {
			return []int{j, i}
		}
		m[nums[i]] = i
	}

	return nil
}
