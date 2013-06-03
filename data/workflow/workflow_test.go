package workflow

import (
	"github.com/bmizerany/assert"
	"playbill/data/stage"
	"testing"
)

func TestNew(t *testing.T) {
	w := New("W0", "")
	assert.Equal(t, len(w.Stages), 0)
}

func TestAddStage(t *testing.T) {
	w := New("W0", "")
	s0 := stage.New("S0", "")
	s1 := stage.New("S1", "")

	// Yep..
	_, err := w.AddStage(s0)
	assert.Equal(t, w.Stages[0], s0)

	// Yep..
	_, err = w.AddStage(s1)
	assert.Equal(t, w.Stages[1], s1)

	// Nope..
	_, err = w.AddStage(s1)
	assert.NotEqual(t, err, nil)
}

func TestInsertStage(t *testing.T) {
	w := New("W0", "")
	s0 := stage.New("S0", "")

	// Cannot insert at a position that does not exist
	_, err := w.InsertStage(0, s0)
	assert.NotEqual(t, err, nil)

	w.AddStage(s0)

	// Insert new stage at position 0
	ns0 := stage.New("New S0", "")
	_, err = w.InsertStage(0, ns0)
	assert.Equal(t, err, nil)
	assert.Equal(t, w.Stages[0], ns0)
	assert.Equal(t, w.Stages[1], s0)
}

func TestMoveStage(t *testing.T) {
	w := New("W0", "")
	s0 := stage.New("S0", "")
	s1 := stage.New("S1", "")

	w.AddStage(s0)
	w.AddStage(s1)

	w.MoveStage(0, s1)
	assert.Equal(t, w.Stages[0], s1)
	assert.Equal(t, w.Stages[1], s0)
}
