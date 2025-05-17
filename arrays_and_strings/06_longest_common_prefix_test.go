package arrays_and_strings

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	require.Equal(t, "fl", longestCommonPrefix([]string{"flower", "flow", "flight"}))

	require.Equal(t, "", longestCommonPrefix([]string{"dog", "racecar", "car"}))

	require.Equal(t, "a", longestCommonPrefix([]string{"a"}))

	require.Equal(t, "c", longestCommonPrefix([]string{"cir", "car"}))

	require.Equal(t, "", longestCommonPrefix([]string{"abab", "aba", ""}))
}
