package main

func moveZeroes(nums []int) {
	stackSize := 0
	for _, x := range nums {
		if x != 0 {
			nums[stackSize] = x
			stackSize++
		}
	}
	clear(nums[stackSize:])
}
