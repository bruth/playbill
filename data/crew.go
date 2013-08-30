package data

import (
	"fmt"
	"regexp"
	"strings"
)

// Matches any characters wrapped in two sets of curly braces. Spaces are
// ignored around the placeholder characters. {{foo}}, {{ foo }}, {{foo }} are
// all equivalent.
var placholderRe = regexp.MustCompile("{{ *([^}]+?) *}}")

type Vars map[string]interface{}
type Env map[string]string

// Render function for trivial template format. Takes a string with {{key}}
func renderTemplate(t string, data map[string]interface{}) (string, error) {
	matches := placholderRe.FindAllStringSubmatch(t, -1)

	for _, match := range matches {
		// raw is the key including the curly braces
		raw, key := match[0], match[1]
		value, ok := data[key]
		// Return error if a key does not exist in the data, but exists in the template
		if !ok {
			return "<error>", fmt.Errorf("no key '%s' exists in data", key)
		}
		// Replace all occurrences of raw with value
		t = strings.Replace(t, raw, fmt.Sprint(value), -1)
	}

	// Replace escaped braces
	t = strings.Replace(t, "\\{", "{", -1)
	t = strings.Replace(t, "\\}", "}", -1)

	return t, nil
}

// Structure that defines environment variables and command variables used during
// a routine's runtime. A `Routine`'s command name and arguments are templates in
// `Vars` defines the template variables. Crews are typically defined per host or
// for isolated environments are a single host.
type Crew struct {
	Name string
	Env  Env
	Vars Vars
	Dir  string
}

// Parses and applies the template variables to a template string
func (c *Crew) Render(t string) (string, error) {
	return renderTemplate(t, c.Vars)
}

// Initializes a new named crew
func NewCrew(n string) *Crew {
	return &Crew{
		Name: n,
	}
}
