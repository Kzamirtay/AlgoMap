package hashmaps_and_sets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanConstruct(t *testing.T) {
	assert.Equal(t, false, canConstruct("a", "b"))

	assert.Equal(t, false, canConstruct("aa", "ab"))

	assert.Equal(t, true, canConstruct("aa", "aab"))

	assert.Equal(t, false, canConstruct("fihjjjjei", "hjibagacbhadfaefdjaeaebgi"))
}
