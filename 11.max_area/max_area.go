// https://leetcode.cn/problems/container-with-most-water/description/?envType=study-plan-v2&envId=top-100-liked
//
// 11. 盛最多水的容器
// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量。
// 说明：你不能倾斜容器。
package maxarea

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	result := 0

	for l < r {
		result = max(result, (r-l)*min(height[r], height[l]))
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}

	return result
}
