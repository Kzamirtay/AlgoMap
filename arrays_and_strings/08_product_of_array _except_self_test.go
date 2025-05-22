package arrays_and_strings

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	require.Equal(t, []int{60, 40, 30, 24}, productExceptSelf([]int{2, 3, 4, 5}))

	require.Equal(t, []int{120, 60, 40, 30, 24}, productExceptSelf([]int{1, 2, 3, 4, 5}))

	require.Equal(t, []int{2, 1}, productExceptSelf([]int{1, 2}))

	require.Equal(t, []int{6, 3, 2}, productExceptSelf([]int{1, 2, 3}))
}
