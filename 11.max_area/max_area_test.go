package maxarea

import "testing"

func Test_maxArea(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		height []int
		want   int
	}{
		{
			name:   "case1",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "case2",
			height: []int{1, 1},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxArea(tt.height)
			if got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
