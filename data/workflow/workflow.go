package workflow

import (
	"errors"
	"playbill/data/stage"
	"time"
)

// A workflow is composed of one or more stages and intends to represent
// a complete cycle of work.
type Workflow struct {
	Name        string
	Description string
	Created     time.Time
	Modified    time.Time
	Stages      []*stage.Stage
}

// Returns true if the stage exists in the workflow
func (w *Workflow) HasStage(s *stage.Stage) bool {
	for _, e := range w.Stages {
		if e == s {
			return true
		}
	}
	return false
}

// Gets the stage by name contained in this workflow
func (w *Workflow) GetStage(name string) (*stage.Stage, error) {
	for _, e := range w.Stages {
		if e.Name == name {
			return e, nil
		}
	}
	return nil, errors.New("stage does not exist in workflow")
}

// Add a stage to the end of the workflow
func (w *Workflow) AddStage(s *stage.Stage) (*Workflow, error) {
	if w.HasStage(s) {
		return nil, errors.New("stage already exists in workflow")
	}
	w.Stages = append(w.Stages, s)
	return w, nil
}

// Insert stage at position `i` in the workflow
func (w *Workflow) InsertStage(i int, s *stage.Stage) (*Workflow, error) {
	if len(w.Stages) <= i || i < 0 {
		return nil, errors.New("cannot insert stage; out of bounds")
	}
	if w.HasStage(s) {
		return nil, errors.New("stage already exists in workflow")
	}
	// Append an element (to create a new underlying array)
	w.Stages = append(w.Stages, nil)
	// Shift elements (first argument is destination)
	copy(w.Stages[i+1:], w.Stages[i:])
	// Set position i
	w.Stages[i] = s
	return w, nil
}

// Remove a stage from the workflow. This currently assumes only
// a stage of each type
func (w *Workflow) RemoveStage(s *stage.Stage) (*Workflow, error) {
	for i, e := range w.Stages {
		if e == s {
			w.Stages = append(w.Stages[:i], w.Stages[i+1:]...)
			return w, nil
		}
	}
	return nil, errors.New("stage does not exist in workflow")
}

// Moves a stage to a new position
func (w *Workflow) MoveStage(i int, s *stage.Stage) (*Workflow, error) {
	if _, err := w.RemoveStage(s); err != nil {
		return nil, err
	}
	if _, err := w.InsertStage(i, s); err != nil {
		return nil, err
	}
	return w, nil
}

// Run the full workflow
func (w *Workflow) Run() (*WorkflowRun, error) {
	r := NewRun(w)
	r.Run()
	return r, nil
}

func New(name string, description string) *Workflow {
	w := &Workflow{
		Name:        name,
		Description: description,
		Created:     time.Now(),
		Modified:    time.Now(),
		Stages:      make([]*stage.Stage, 0),
	}
	return w
}
