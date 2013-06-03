package stage

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestSuccessfulRun(t *testing.T) {
	s := New("S0", "")

	s.NewTask("echo", "Hello")
	s.NewTask("echo", "Gopher")
	s.NewTask("echo", "World")

	sr, err := s.Run()

	assert.Equal(t, err, nil)
	assert.Equal(t, len(sr.Tasks), 3)
	assert.Equal(t, sr.Failed, false)
	assert.Equal(t, sr.Done(), true)
}

func TestFailedRun(t *testing.T) {
	s := New("S0", "")

	s.NewTask("git", "flub")
	s.NewTask("echo", "Gopher")
	s.NewTask("echo", "World")

	sr, err := s.Run()

	assert.Equal(t, err, nil)
	assert.Equal(t, len(sr.Tasks), 3)
	assert.Equal(t, sr.Failed, true)
	assert.Equal(t, sr.Done(), true)
}
