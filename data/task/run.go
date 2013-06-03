package task

import (
	"time"
)

// Logs various data about the execution of the task
type TaskRun struct {
	Task *Task

	// Start and end timestamps
	StartTime time.Time
	EndTime   time.Time

	// Flags denoting the state of the run
	Failed  bool
	Skipped bool
	Aborted bool

	// Task command interfaces
	Stdin  string
	Stdout string
	Stderr string
}

func (r *TaskRun) Done() bool {
	return !r.EndTime.IsZero()
}

func NewRun(t *Task) *TaskRun {
	return &TaskRun{
		Task: t,
	}
}
