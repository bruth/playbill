package data

// Structure that defines environment variables and command variables used during
// a routine's runtime. A `Routine`'s command name and arguments are templates in
// `Vars` defines the template variables. Crews are typically defined per host or
// for isolated environments are a single host.
type Crew struct {
	Name string                 `bson:"name" json:"name"`
	Env  []string               `bson:"env" json:"env"`
	Vars map[string]interface{} `bson:"vars" json:"vars"`
	Dir  string                 `bson:"dir" json:"dir"`
}

// Parses and applies the template variables to a template string
func (c *Crew) Render(t string) (string, error) {
	return renderTemplate(t, c.Vars)
}

// Initializes a new named crew
func NewCrew(n string) *Crew {
	return &Crew{
		Name: n,
		Env:  make([]string, 0),
		Vars: map[string]interface{}{},
	}
}

// Initializes a new crew from a file
func ImportCrew(path string) (*Crew, error) {
    var c Crew
    err := ImportComponent(path, &c)
    return &c, err
}
