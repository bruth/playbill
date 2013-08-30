package data

import (
	"errors"
	"fmt"
	"time"
)

var ErrSceneFailed = errors.New("Scene failed")

// Logs the start and end time as well as whether the play
// has been aborted.
type PlayRun struct {
	Play      *Play
	Crew      *Crew
	Scenes    []*SceneRun
	StartTime time.Time
	EndTime   time.Time
	Failed    bool
	Skipped   bool
	Aborted   bool
}

func (r *PlayRun) Done() bool {
	return !r.EndTime.IsZero()
}

// Starts the run
func (r *PlayRun) Run() error {
	p := r.Play
	r.StartTime = time.Now()

	// Run each scene, if any scene fails, mark the play run as failed
	// and return
	for i, s := range p.Scenes {
		sr := NewSceneRun(s, r.Crew)
		r.Scenes[i] = sr
		if err := sr.Run(); err != nil {
			r.Failed = true
			fmt.Println(err)
			break
		}
	}

	r.EndTime = time.Now()

	if r.Failed {
		return ErrSceneFailed
	}
	return nil
}

func NewPlayRun(p *Play, c *Crew) *PlayRun {
	return &PlayRun{
		Play:   p,
		Crew:   c,
		Scenes: make([]*SceneRun, len(p.Scenes)),
	}
}
