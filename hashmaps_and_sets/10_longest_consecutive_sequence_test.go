package hashmaps_and_sets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	assert.Equal(t, 4, longestConsecutive([]int{100, 4, 200, 1, 3, 2}))

	assert.Equal(t, 9, longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))

	assert.Equal(t, 3, longestConsecutive([]int{1, 0, 1, 2}))

	assert.Equal(t, 7, longestConsecutive([]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}))
}
