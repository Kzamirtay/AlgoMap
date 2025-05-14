package arrays_and_strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindClosestNumber(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1}
	assert.Equal(t, 1, findClosestNumber(nums))

	nums = []int{-4, -2, 1, 4, 8}
	assert.Equal(t, 1, findClosestNumber(nums))

	nums = []int{2, -1, 1}
	assert.Equal(t, 1, findClosestNumber(nums))

	nums = []int{-10, 10, 10, 10}
	assert.Equal(t, 10, findClosestNumber(nums))

	nums = []int{-1, 2}
	assert.Equal(t, -1, findClosestNumber(nums))

	nums = []int{-100000, -100000}
	assert.Equal(t, -100000, findClosestNumber(nums))

	nums = []int{-10, -12, -54, -12, -544, -10000}
	assert.Equal(t, -10, findClosestNumber(nums))
}

func TestAbs(t *testing.T) {
	assert.Equal(t, 1, abs(1))
	assert.Equal(t, 123, abs(123))
	assert.Equal(t, 1, abs(-1))
	assert.Equal(t, 123, abs(-123))
	assert.Equal(t, 0, abs(0))
}
