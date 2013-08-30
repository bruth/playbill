package data

import (
	"github.com/bruth/assert"
	"testing"
)

func TestSuccessfulRoutineRun(t *testing.T) {
	r := NewRoutine("echo", "hello test")
	c := NewCrew("c1")
	rr, err := r.Run(c)

	assert.Equal(t, err, nil)
	assert.Equal(t, rr.Done(), true)
	assert.Equal(t, rr.Failed, false)
	assert.Equal(t, rr.Stdin, "")
	assert.Equal(t, rr.Stdout, "hello test\n")
	assert.Equal(t, rr.Stderr, "")
}

func TestFailedRoutineRun(t *testing.T) {
	r := NewRoutine("git", "flub")
	c := NewCrew("c1")
	rr, err := r.Run(c)

	assert.NotEqual(t, err, nil)
	assert.Equal(t, rr.Done(), true)
	assert.Equal(t, rr.Failed, true)
	assert.Equal(t, rr.Stdin, "")
	assert.Equal(t, rr.Stdout, "")
	assert.Equal(t, rr.Stderr, "git: 'flub' is not a git command. See 'git --help'.\n\nDid you mean this?\n\tpull\n")
}
