package data

import (
	"errors"
	"time"
)

var (
	ErrRoutineExists   = errors.New("Routine already exists for scene")
	ErrNoRoutineExists = errors.New("Routine does not exist for scene")
)

// A single scene of a play. A scene is composed of one or routines
// that are required to be executed for this scene to be considered
// complete. Note that all routines for a scene are executed in parallel.
// The assumption exists, since, if a routine is serially processed,
// it should be defined as it's own scene.
type Scene struct {
	Name        string     `bson:"name" json:"name"`
	Description string     `bson:"description" json:"description"`
	Created     time.Time  `bson:"created" json:"created"`
	Modified    time.Time  `bson:"modified" json:"modified"`
	Routines    []*Routine `bson:"routines" json:"routines"`
}

// Creates a new routines and adds it to the scene
func (s *Scene) NewRoutine(name string, args ...string) *Routine {
	r := NewRoutine(name, args...)
	s.AddRoutine(r)
	return r
}

// Adds an existing routine to the scene
func (s *Scene) AddRoutine(r *Routine) error {
	if s.HasRoutine(r) {
		return ErrRoutineExists
	}
	s.Routines = append(s.Routines, r)
	return nil
}

// Checks if a routine exists in the scene
func (s *Scene) HasRoutine(r *Routine) bool {
	for _, e := range s.Routines {
		if e == r {
			return true
		}
	}
	return false
}

// Gets the routine by name contained in this scene
func (s *Scene) GetRoutine(name string) (*Routine, error) {
	for _, e := range s.Routines {
		if e.Name == name {
			return e, nil
		}
	}
	return nil, ErrNoRoutineExists
}

// Runs routine's command and redirects all standard output to the OS
func (s *Scene) Run(c *Crew) (*SceneRun, error) {
	r := NewSceneRun(s, c)
	r.Run()
	return r, nil
}

func NewScene(name string, description string) *Scene {
	s := &Scene{
		Name:        name,
		Description: description,
		Created:     time.Now(),
		Modified:    time.Now(),
		Routines:    make([]*Routine, 0),
	}
	return s
}
