package longestconsecutive

// 1. 排序后指针
// 理论上时间复杂度 O(NlogN)
// 但在leetcode运行比方法二的O(N)快
//
// func longestConsecutive(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
//
// 	slices.Sort(nums)
//
// 	j := 1
// 	tmp := 1
// 	res := 1
// 	for j < len(nums) {
// 		if nums[j-1] == nums[j] {
// 			j++
// 			continue
// 		}
//
// 		if nums[j-1] == nums[j]-1 {
// 			j++
// 			tmp++
// 			res = max(res, tmp)
// 			continue
// 		}
//
// 		tmp = 1
// 		j = j + 1
// 	}
//
// 	return res
// }

// 2. 遍历生成map，然后遍历map
// 时间复杂度 O(N)
func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})

	for _, v := range nums {
		m[v] = struct{}{}
	}

	res := 0
	for v := range m {
		if _, ok := m[v-1]; ok {
			continue
		}

		tmp := 1
		for {
			if _, ok := m[v+1]; !ok {
				res = max(res, tmp)
				break
			}
			v++
			tmp++
		}
	}

	return res
}
