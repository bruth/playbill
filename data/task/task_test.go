package task

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestSuccessfulRun(t *testing.T) {
	tsk := New("echo", "hello test")
	tr, err := tsk.Run()

	assert.Equal(t, err, nil)
	assert.Equal(t, tr.Done(), true)
	assert.Equal(t, tr.Failed, false)
	assert.Equal(t, tr.Stdin, "")
	assert.Equal(t, tr.Stdout, "hello test\n")
	assert.Equal(t, tr.Stderr, "")
}

func TestFailedRun(t *testing.T) {
	tsk := New("git", "flub")
	tr, err := tsk.Run()

	assert.NotEqual(t, err, nil)
	assert.Equal(t, tr.Done(), true)
	assert.Equal(t, tr.Failed, true)
	assert.Equal(t, tr.Stdin, "")
	assert.Equal(t, tr.Stdout, "")
	assert.Equal(t, tr.Stderr, "git: 'flub' is not a git command. See 'git --help'.\n\nDid you mean this?\n\tpull\n")
}
