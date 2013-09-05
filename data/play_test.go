package data

import (
	"github.com/bruth/assert"
	"testing"
)

func TestNewPlay(t *testing.T) {
	p := NewPlay("P0", "")
	assert.Equal(t, len(p.Scenes), 0)
}

func TestPlayAddScene(t *testing.T) {
	p := NewPlay("P0", "")
	s0 := NewScene("S0", "")
	s1 := NewScene("S1", "")

	// Yep..
	p.AddScene(s0)
	assert.Equal(t, p.Scenes[0], s0)

	// Yep..
	p.AddScene(s1)
	assert.Equal(t, p.Scenes[1], s1)

	// Nope..
	err := p.AddScene(s1)
	assert.NotEqual(t, err, nil)
}

func TestPlayInsertScene(t *testing.T) {
	p := NewPlay("P0", "")
	s0 := NewScene("S0", "")

	// Cannot insert at a position that does not exist
	err := p.InsertScene(0, s0)
	assert.NotEqual(t, err, nil)

	p.AddScene(s0)

	// Insert new scene at position 0
	ns0 := NewScene("New S0", "")
	err = p.InsertScene(0, ns0)
	assert.Equal(t, err, nil)
	assert.Equal(t, p.Scenes[0], ns0)
	assert.Equal(t, p.Scenes[1], s0)
}

func TestPlayMoveScene(t *testing.T) {
	p := NewPlay("P0", "")
	s0 := NewScene("S0", "")
	s1 := NewScene("S1", "")

	p.AddScene(s0)
	p.AddScene(s1)

	p.MoveScene(0, s1)
	assert.Equal(t, p.Scenes[0], s1)
	assert.Equal(t, p.Scenes[1], s0)
}

func TestPlayRun(t *testing.T) {
	p := NewPlay("P0", "")
	c := &Crew{}

	// Nothing to do..
	pr, err := p.Run(c)
	assert.Equal(t, err, nil)
	assert.False(t, pr.Failed)

	// Scene and routine
	s0, _ := p.NewScene("S0", "")
	s0.NewRoutine("echo", "foo")
	pr, err = p.Run(c)
	assert.Equal(t, err, nil)
	assert.False(t, pr.Failed)
}
