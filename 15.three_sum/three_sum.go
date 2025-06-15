// https://leetcode.cn/problems/3sum/description/?envType=study-plan-v2&envId=top-100-liked
//
// 15. 三数之和
// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。
package main

import (
	"fmt"
	"slices"
)

func main() {
	// res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	res := threeSum([]int{0, 0, 0, 0})
	fmt.Println(res)
}

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	res := [][]int{}
	length := len(nums)

	for i := 0; i <= length-3; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, length-1
		for l < r {
			sum := nums[l] + nums[r] + nums[i]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return res
}
