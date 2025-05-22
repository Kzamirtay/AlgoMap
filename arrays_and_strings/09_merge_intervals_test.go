package arrays_and_strings

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	require.Equal(t, [][]int{
		{1, 6},
		{8, 10},
		{15, 18},
	}, merge([][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}))

	require.Equal(t, [][]int{
		{1, 5},
	}, merge([][]int{
		{1, 4},
		{4, 5},
	}))

	require.Equal(t, [][]int{
		{1, 4},
		{5, 6},
	}, merge([][]int{
		{1, 4},
		{5, 6},
	}))

	require.Equal(t, [][]int{
		{0, 4},
	}, merge([][]int{
		{1, 4},
		{0, 4},
	}))

	require.Equal(t, [][]int{
		{1, 4},
	}, merge([][]int{
		{1, 4},
		{2, 3},
	}))
}
