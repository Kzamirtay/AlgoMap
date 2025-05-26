package hashmaps_and_sets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	assert.Equal(t, true, isAnagram("anagram", "nagaram"))

	assert.Equal(t, false, isAnagram("rat", "car"))
}
