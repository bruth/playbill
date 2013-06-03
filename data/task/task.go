package task

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
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

func startTask(r *TaskRun) error {
	r.StartTime = time.Now()
	if r.Task.CmdName == "" {
		return endTask(r, ErrNoCommand)
	}
	return nil
}

func endTask(r *TaskRun, err error) error {
	r.EndTime = time.Now()
	r.Failed = err != nil
	return err
}

// Internal constructor which takes an additional argument for
// tee interfaces to OS standard in/out/err in addition to writing
// it to the struct
func runTask(r *TaskRun, store bool, tee bool) error {
	startTask(r)

	// Construct the command to be executed
	c := exec.Command(r.Task.CmdName, r.Task.CmdArgs...)

	// Buffers for subcommand interfaces. These are necessary for storing
	// on to the task run. Note, these may not actually be used below.
	rstdin := new(bytes.Buffer)
	rstdout := new(bytes.Buffer)
	rstderr := new(bytes.Buffer)

	if store && tee {
		// Reads from os.Stdin and sends it to c.Stdin as well as writing
		// it to the stdin buffer
		c.Stdin = io.TeeReader(os.Stdin, rstdin)
		// Writes the commands stdout to os.Stdout and the buffer
		c.Stdout = io.MultiWriter(os.Stdout, rstdout)
		c.Stderr = io.MultiWriter(os.Stderr, rstderr)
	} else if store {
		// stdin must always come from the OS
		c.Stdin = io.TeeReader(os.Stdin, rstdin)
		c.Stdout = rstdout
		c.Stderr = rstderr
	} else if tee {
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
	}

	// Run command
	fmt.Printf("Running task `%s`\n", r.Task.CmdString())
	err := c.Run()

	if store {
		// Set flags, store std interfaces
		r.Stdin = rstdin.String()
		r.Stdout = rstdout.String()
		r.Stderr = rstderr.String()
	}

	endTask(r, err)

	return err
}

// Run the task's command
func (r *TaskRun) Run() error {
	return runTask(r, true, false)
}

// Run the tasks's command in 'debug' mode which redirects the command
// std interfaces to the OS
func (r *TaskRun) RunDebug() error {
	return runTask(r, false, true)
}

// Run the task's command and tees all command std interfaces to OS
func (r *TaskRun) RunTee() error {
	return runTask(r, true, true)
}
