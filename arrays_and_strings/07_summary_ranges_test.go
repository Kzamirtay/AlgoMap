package arrays_and_strings

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	require.Equal(t, []string(nil), summaryRanges([]int{}))

	assert.Equal(t, []string{"0"}, summaryRanges([]int{0}))

	assert.Equal(t, []string{"0->3"}, summaryRanges([]int{0, 1, 2, 3}))

	assert.Equal(t, []string{"0->2", "4->5"}, summaryRanges([]int{0, 1, 2, 4, 5}))

	assert.Equal(t, []string{"0->2", "4->5", "7"}, summaryRanges([]int{0, 1, 2, 4, 5, 7}))

	assert.Equal(t, []string{"0", "2->4", "6", "8->9"}, summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}
