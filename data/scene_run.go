package data

import (
	"errors"
	"log"
	"time"
)

var ErrRoutineFailed = errors.New("Routine failed")

// Logs the start and end time as well as whether the scene has
// been aborted
type SceneRun struct {
	Scene     *Scene
	Crew      *Crew
	Routines  []*RoutineRun
	StartTime time.Time
	EndTime   time.Time
	Failed    bool
	Skipped   bool
	Aborted   bool
}

func (r *SceneRun) Done() bool {
	return !r.EndTime.IsZero()
}

// Runs a routine and sends the error into the channel. This is intended
// to be run in a separate goroutine
func runSceneRoutine(r *RoutineRun, c chan error) {
	err := r.Run()
	c <- err
}

// Runs the scenes
func (sr *SceneRun) Run() error {
	s := sr.Scene

	// Initialize a new channel with at most N slots
	c := make(chan error, len(s.Routines))

	// Create new runs for each routine and execute them in separate
	// goroutines
	for i, r := range s.Routines {
		_r := r.NewRun(sr.Crew)
		sr.Routines[i] = _r
		go runSceneRoutine(_r, c)
	}

	// Block and receive errors from channel. Set scene as failed
	// if any are present and print them
	for _ = range s.Routines {
		if err := <-c; err != nil {
			sr.Failed = true
			log.Println(err)
		}
	}

	sr.EndTime = time.Now()

	if sr.Failed {
		return ErrRoutineFailed
	}
	return nil
}

func NewSceneRun(s *Scene, c *Crew) *SceneRun {
	return &SceneRun{
		Scene:    s,
		Crew:     c,
		Routines: make([]*RoutineRun, len(s.Routines)),
	}
}
