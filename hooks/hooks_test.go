package hooks

import (
	"github.com/bruth/assert"
	"testing"
)

func TestParseHooksFile(t *testing.T) {
	assert.True(t, true)
}

func TestTrigger(t *testing.T) {
	// Null case
	c, _ := Trigger("testhook", nil)
	assert.Equal(t, c, 0)

	// Local file
	NewLocalHooks()
}
