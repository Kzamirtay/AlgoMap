package arrays_and_strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeAlternately(t *testing.T) {
	assert.Equal(t, "ap", mergeAlternately("a", "p"))

	assert.Equal(t, "apbqcr", mergeAlternately("abc", "pqr"))

	assert.Equal(t, "apbqrs", mergeAlternately("ab", "pqrs"))

	assert.Equal(t, "apbqcd", mergeAlternately("abcd", "pq"))
}
