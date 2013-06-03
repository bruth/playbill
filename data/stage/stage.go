package stage

import (
	"errors"
	"playbill/data/task"
	"time"
)

var (
	ErrTaskExists   = errors.New("Task already exists for stage")
	ErrNoTaskExists = errors.New("Task does not exist for stage")
)

// A single stage of a workflow. A stage is composed of one or tasks
// that are required to be executed for this stage to be considered
// complete. Note that all tasks for a stage are executed in parallel.
// The assumption exists, since, if a task is serially processed,
// it should be defined as it's own stage.
type Stage struct {
	Name        string
	Description string
	Created     time.Time
	Modified    time.Time
	Tasks       []*task.Task
}

// Creates a new tasks and adds it to the stage
func (s *Stage) NewTask(name string, args ...string) *task.Task {
	t := task.New(name, args...)
	s.AddTask(t)
	return t
}

// Adds an existing task to the stage
func (s *Stage) AddTask(t *task.Task) (*Stage, error) {
	if s.HasTask(t) {
		return nil, ErrTaskExists
	}
	s.Tasks = append(s.Tasks, t)
	return s, nil
}

// Checks if a task exists in the stage
func (s *Stage) HasTask(t *task.Task) bool {
	for _, e := range s.Tasks {
		if e == t {
			return true
		}
	}
	return false
}

// Gets the task by name contained in this stage
func (s *Stage) GetTask(name string) (*task.Task, error) {
	for _, e := range s.Tasks {
		if e.Name == name {
			return e, nil
		}
	}
	return nil, ErrNoTaskExists
}

// Runs task's command and redirects all standard output to the OS
func (s *Stage) Run() (*StageRun, error) {
	r := NewRun(s)
	r.Run()
	return r, nil
}

func New(name string, description string) *Stage {
	s := &Stage{
		Name:        name,
		Description: description,
		Created:     time.Now(),
		Modified:    time.Now(),
		Tasks:       make([]*task.Task, 0),
	}
	return s
}
