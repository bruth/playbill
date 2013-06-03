package workflow

import (
	"errors"
	"fmt"
	"playbill/data/stage"
	"time"
)

var (
	ErrorStageFailed = errors.New("Stage failed")
)

// Logs the start and end time as well as whether the workflow
// has been aborted.
type WorkflowRun struct {
	Workflow  *Workflow
	Stages    []*stage.StageRun
	StartTime time.Time
	EndTime   time.Time
	Failed    bool
	Skipped   bool
	Aborted   bool
}

func (r *WorkflowRun) Done() bool {
	return !r.EndTime.IsZero()
}

func (r *WorkflowRun) Run() error {
	w := r.Workflow
	r.StartTime = time.Now()

	for i, s := range w.Stages {
		sr := stage.NewRun(s)
		r.Stages[i] = sr
		if err := sr.Run(); err != nil {
			r.Failed = true
			fmt.Println(err)
		}
	}

	r.EndTime = time.Now()

	if r.Failed {
		return ErrorStageFailed
	}
	return nil
}

func NewRun(w *Workflow) *WorkflowRun {
	return &WorkflowRun{
		Workflow: w,
		Stages:   make([]*stage.StageRun, len(w.Stages)),
	}
}
