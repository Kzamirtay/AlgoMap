package hashmaps_and_sets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumJewelsInStones(t *testing.T) {
	assert.Equal(t, 3, numJewelsInStones("aA", "aAAbbbb"))

	assert.Equal(t, 0, numJewelsInStones("z", "ZZ"))
}
