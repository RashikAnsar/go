package twonumbersum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoNumberSum(t *testing.T) {
	testCases := []struct {
		nums     []int
		goal     int
		expected []int
	}{
		{
			nums:     []int{1, 9, 13, 20, 47},
			goal:     10,
			expected: []int{1, 9},
		},
		{
			nums:     []int{3, 5, -4, 8, 11, 1, -1, 6},
			goal:     10,
			expected: []int{-1, 11},
		},
		{
			nums:     []int{3, 2, 4, 1, 9},
			goal:     12,
			expected: []int{3, 9},
		},
		{
			nums:     []int{-7, -5, -3, -1, 0, 1, 3, 5, 7},
			goal:     -5,
			expected: []int{-5, 0},
		},
		{
			nums:     []int{},
			goal:     10,
			expected: []int{},
		},
	}
	for idx, tC := range testCases {
		t.Run(fmt.Sprintf("test %d - TwoNumberSum should return expected output", idx), func(t *testing.T) {
			output := TwoNumberSum(tC.nums, tC.goal)
			require.ElementsMatch(t, tC.expected, output)
		})
	}
}
