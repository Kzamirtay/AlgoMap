package arrays_and_strings

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	require.Equal(t, 5, maxProfit([]int{7, 1, 5, 3, 6, 4}))

	require.Equal(t, 0, maxProfit([]int{7, 6, 4, 3, 1}))

	require.Equal(t, 2, maxProfit([]int{2, 4, 1}))

	require.Equal(t, 2, maxProfit([]int{7, 2, 4, 1}))

	require.Equal(t, 0, maxProfit([]int{1}))
}
