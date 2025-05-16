package arrays_and_strings

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	require.Equal(t, true, isSubsequence("abc", "ahbgdc"))

	require.Equal(t, false, isSubsequence("axc", "ahbgdc"))

	require.Equal(t, false, isSubsequence("acb", "ahbgdc"))

	require.Equal(t, false, isSubsequence("aaaaaa", "bbaaaa"))

	require.Equal(t, true, isSubsequence("ab", "baab"))
}
