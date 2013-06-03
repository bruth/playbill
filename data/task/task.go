package task

import (
	"errors"
	"fmt"
	"strings"
)

var ErrNoCommand = errors.New("Command not defined")

// A task represents a container for an external command
type Task struct {
	// Name of the task
	Name string

	// Description of the task
	Description string

	// Command name
	CmdName string

	// Array of command line arguments
	CmdArgs []string
}

// Returns a string of the command
func (t *Task) CmdString() string {
	if len(t.CmdArgs) == 0 {
		return t.CmdName
	}
	return fmt.Sprintf("%s %s", t.CmdName, strings.Join(t.CmdArgs, " "))
}

// Creates a run for the task
func (t *Task) NewRun() *TaskRun {
	return &TaskRun{
		Task: t,
	}
}

// Shortcut for creating and immediately running the task
func (t *Task) Run() (*TaskRun, error) {
	r := t.NewRun()
	err := r.Run()
	return r, err
}

// Creates a new task
func New(name string, args ...string) *Task {
	t := &Task{
		CmdName: name,
		CmdArgs: args,
	}
	return t
}
