package data

import (
	"github.com/bruth/assert"
	"testing"
)

func TestSuccessfulSceneRun(t *testing.T) {
	s := NewScene("S0", "")

	s.NewRoutine("echo", "Hello")
	s.NewRoutine("echo", "Gopher")
	s.NewRoutine("echo", "World")

	c := NewCrew("c1")
	sr, err := s.Run(c)

	assert.Equal(t, err, nil)
	assert.Equal(t, len(sr.Routines), 3)
	assert.Equal(t, sr.Failed, false)
	assert.Equal(t, sr.Done(), true)
}

func TestFailedSceneRun(t *testing.T) {
	s := NewScene("S0", "")

	s.NewRoutine("git", "flub")
	s.NewRoutine("echo", "Gopher")
	s.NewRoutine("echo", "World")

	c := NewCrew("c1")
	sr, err := s.Run(c)

	assert.Equal(t, err, nil)
	assert.Equal(t, len(sr.Routines), 3)
	assert.Equal(t, sr.Failed, true)
	assert.Equal(t, sr.Done(), true)
}
