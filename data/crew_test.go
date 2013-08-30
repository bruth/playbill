package data

import (
	"github.com/bruth/assert"
	"testing"
)

func TestCrewRenderPlain(t *testing.T) {
	c := NewCrew("test")

	// Set some variables
	c.Vars = map[string]interface{}{
		"path": "/usr/local/bin",
		"n":    4,
		"x":    3.45,
	}

	// No placeholders
	s, _ := c.Render("python test.py")
	assert.Equal(t, s, "python test.py")

	// Missing key
	s, err := c.Render("{{bindir}}/python test.py")
	assert.NotEqual(t, err, nil)

	// Typical
	s, err = c.Render("{{path}}/python -n {{n}} -x {{x}} test.py")
	assert.Equal(t, s, "/usr/local/bin/python -n 4 -x 3.45 test.py")

	// Single braces aren't affected
	s, err = c.Render("{path}")
	assert.Equal(t, s, "{path}")

	// Escaped braces aren't affected, but replaced
	s, err = c.Render("{\\{path\\}}")
	assert.Equal(t, s, "{{path}}")

	// Escaped braces, round 2
	s, err = c.Render("\\{\\{path\\}\\}")
	assert.Equal(t, s, "{{path}}")
}
