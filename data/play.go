package data

import (
	"errors"
	"time"
)

var (
	ErrSceneExists      = errors.New("Scene already exists in play")
	ErrNoSceneExists    = errors.New("Scene does not exist in play")
	ErrSceneOutOfBounds = errors.New("Cannot insert scene; out of bounds")
)

// A play is composed of one or more scenes and intends to represent
// a complete cycle of work.
type Play struct {
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
	Created     time.Time     `bson:"created" json:"created"`
	Modified    time.Time     `bson:"modified" json:"modified"`
	Scenes      []*Scene      `bson:"scenes" json:"scenes"`
}

// Create a new scene and add it to the play
func (p *Play) NewScene(name string, description string) (*Scene, error) {
	s := NewScene(name, description)
	err := p.AddScene(s)
	return s, err
}

// Returns true if the scene exists in the play
func (p *Play) HasScene(s *Scene) bool {
	for _, e := range p.Scenes {
		if e == s {
			return true
		}
	}
	return false
}

// Gets the scene by name contained in this play
func (p *Play) GetScene(name string) (*Scene, error) {
	for _, s := range p.Scenes {
		if s.Name == name {
			return s, nil
		}
	}
	return nil, ErrSceneExists
}

// Add a scene to the end of the play
func (p *Play) AddScene(s *Scene) error {
	if p.HasScene(s) {
		return ErrSceneExists
	}
	p.Scenes = append(p.Scenes, s)
	return nil
}

// Insert scene at position `i` in the play
func (p *Play) InsertScene(i int, s *Scene) error {
	if i < 0 || len(p.Scenes) <= i {
		return ErrSceneOutOfBounds
	}
	if p.HasScene(s) {
		return ErrSceneExists
	}
	// Append an element (to create a new underlying array)
	p.Scenes = append(p.Scenes, nil)
	// Shift elements (first argument is destination)
	copy(p.Scenes[i+1:], p.Scenes[i:])
	// Set position i
	p.Scenes[i] = s
	return nil
}

// Remove a scene from the play. This currently assumes only
// a scene of each type
func (p *Play) RemoveScene(s *Scene) error {
	for i, e := range p.Scenes {
		if e == s {
			p.Scenes = append(p.Scenes[:i], p.Scenes[i+1:]...)
			return nil
		}
	}
	return ErrNoSceneExists
}

// Moves a scene to a new position
func (p *Play) MoveScene(i int, s *Scene) error {
	if err := p.RemoveScene(s); err != nil {
		return err
	}
	if err := p.InsertScene(i, s); err != nil {
		return err
	}
	return nil
}

// Run the play
func (p *Play) Run(c *Crew) (*PlayRun, error) {
	r := NewPlayRun(p, c)
	r.Run()
	return r, nil
}

func NewPlay(name string, description string) *Play {
	p := &Play{
		Name:        name,
		Description: description,
		Created:     time.Now(),
		Modified:    time.Now(),
		Scenes:      make([]*Scene, 0),
	}
	return p
}
