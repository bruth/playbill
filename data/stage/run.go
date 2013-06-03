package stage

import (
	"errors"
	"fmt"
	"playbill/data/task"
	"time"
)

var (
	ErrTaskFailed = errors.New("Task failed")
)

// Logs the start and end time as well as whether the stage has
// been aborted
type StageRun struct {
	Stage     *Stage
	Tasks     []*task.TaskRun
	StartTime time.Time
	EndTime   time.Time
	Failed    bool
	Skipped   bool
	Aborted   bool
}

func (r *StageRun) Done() bool {
	return !r.EndTime.IsZero()
}

// Runs a task and sends the error into the channel. This is intended
// to be run in a separate goroutine
func runTask(r *task.TaskRun, c chan error) {
	err := r.Run()
	c <- err
}

// Runs the stages
func (r *StageRun) Run() error {
	s := r.Stage

	// Initialize a new channel with at most N slots
	c := make(chan error, len(s.Tasks))

	// Create new runs for each task and execute them in separate
	// goroutines
	for i, t := range s.Tasks {
		tr := t.NewRun()
		r.Tasks[i] = tr
		go runTask(tr, c)
	}

	// Block and receive errors from channel. Set stage as failed
	// if any are present and print them
	for _ = range s.Tasks {
		if err := <-c; err != nil {
			r.Failed = true
			fmt.Println(err)
		}
	}

	r.EndTime = time.Now()

	if r.Failed {
		return ErrTaskFailed
	}
	return nil
}

func NewRun(s *Stage) *StageRun {
	return &StageRun{
		Stage: s,
		Tasks: make([]*task.TaskRun, len(s.Tasks)),
	}
}
