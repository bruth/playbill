package data

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Logs various data about the execution of the routine
type RoutineRun struct {
	Routine *Routine
	Crew    *Crew

	// Start and end timestamps
	StartTime time.Time
	EndTime   time.Time

	// Flags denoting the state of the run
	Failed  bool
	Skipped bool
	Aborted bool

	// Routine command interfaces
	Stdin  string
	Stdout string
	Stderr string
}

func commandParts(r *RoutineRun) (string, []string, error) {
	name, err := r.Crew.Render(r.Routine.Command)
	if err != nil {
		return name, nil, err
	}

	args := make([]string, len(r.Routine.Args))

	for i, arg := range r.Routine.Args {
		arg, err := r.Crew.Render(arg)
		if err != nil {
			return name, nil, err
		}
		args[i] = arg
	}
	return name, args, err
}

func (r *RoutineRun) Command() (*exec.Cmd, error) {
	name, args, err := commandParts(r)
	if err != nil {
		return nil, err
	}
	c := exec.Command(name, args...)
	c.Env = r.Crew.Env
	c.Dir = r.Crew.Dir
	return c, nil
}

func (r *RoutineRun) CommandString() string {
	name, args, err := commandParts(r)
	if err != nil {
		return "<error>"
	}
	pargs := make([]string, len(args)+1)
	pargs[0] = name
	copy(pargs[1:], args)
	return strings.Join(pargs, " ")
}

// Run the routine's command
func (r *RoutineRun) Run() error {
	return runRoutine(r, true, false)
}

// Run the routines's command in 'debug' mode which redirects the command
// std interfaces to the OS
func (r *RoutineRun) RunDebug() error {
	return runRoutine(r, false, true)
}

// Run the routine's command and tees all command std interfaces to OS
func (r *RoutineRun) RunTee() error {
	return runRoutine(r, true, true)
}

func (r *RoutineRun) Done() bool {
	return !r.EndTime.IsZero()
}

func NewRoutineRun(r *Routine, c *Crew) *RoutineRun {
	return &RoutineRun{
		Routine: r,
		Crew:    c,
	}
}

func startRoutine(r *RoutineRun) error {
	r.StartTime = time.Now()

	if r.Routine.Command == "" {
		return endRoutine(r, ErrNoCommand)
	}
	return nil
}

func endRoutine(r *RoutineRun, err error) error {
	r.EndTime = time.Now()
	r.Failed = err != nil
	return err
}

// Internal constructor which takes an additional argument for
// tee interfaces to OS standard in/out/err in addition to writing
// it to the struct
func runRoutine(r *RoutineRun, store bool, tee bool) error {
	startRoutine(r)

	// Construct the command to be executed
	c, err := r.Command()

	if err != nil {
		endRoutine(r, err)
		return err
	}

	// Buffers for subcommand interfaces. These are necessary for storing
	// on to the routine run. Note, these may not actually be used below.
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
	log.Printf("Running routine `%s`\n", r.CommandString())
	err = c.Run()

	if store {
		// Set flags, store std interfaces
		r.Stdin = rstdin.String()
		r.Stdout = rstdout.String()
		r.Stderr = rstderr.String()
	}

	endRoutine(r, err)

	return err
}
