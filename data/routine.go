package data

import (
	"errors"
)

var ErrNoCommand = errors.New("Command not defined")

// A routine represents a container for an external command
type Routine struct {
	Name        string   `bson:"name" json:"name"`
	Description string   `bson:"description" json:"description"`
	Command     string   `bson:"command" json:"command"`
	Args        []string `bson:"args" json:"args"`
}

// Creates a run for the routine
func (r *Routine) NewRun(c *Crew) *RoutineRun {
	return &RoutineRun{
		Routine: r,
		Crew:    c,
	}
}

// Shortcut for creating and immediately running the routine
func (r *Routine) Run(c *Crew) (*RoutineRun, error) {
	_r := r.NewRun(c)
	err := _r.Run()
	return _r, err
}

// Creates a new routine
func NewRoutine(name string, args ...string) *Routine {
	r := &Routine{
		Command: name,
		Args:    args,
	}
	return r
}
