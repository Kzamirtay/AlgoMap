package hashmaps_and_sets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxNumberOfBalloons(t *testing.T) {
	assert.Equal(t, 1, maxNumberOfBalloons("nlaebolko"))

	assert.Equal(t, 2, maxNumberOfBalloons("loonbalxballpoon"))

	assert.Equal(t, 0, maxNumberOfBalloons("leetcode"))
}
