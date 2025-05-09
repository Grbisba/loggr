package loggr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApply(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		var f optionFunc

		assert.NoError(t, f.apply(nil))
	})
}
