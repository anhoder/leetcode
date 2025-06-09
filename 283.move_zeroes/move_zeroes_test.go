package main

import (
	"slices"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want []int
	}{
		{
			name: "case1",
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "case2",
			nums: []int{0},
			want: []int{0},
		},
		{
			name: "case3",
			nums: []int{1, 0},
			want: []int{1, 0},
		},
		{
			name: "case4",
			nums: []int{1, 0, 1},
			want: []int{1, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.nums)
			if !slices.Equal(tt.nums, tt.want) {
				t.Errorf("moveZeroes() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
