package data

import (
	"fmt"
	"regexp"
	"strings"
)

// Matches any characters wrapped in two sets of curly braces. Spaces are
// ignored around the placeholder characters. {{foo}}, {{ foo }}, {{foo }} are
// all equivalent.
var templateVarRe = regexp.MustCompile("{{ *([^}]+?) *}}")

// Render function for trivial template format. Takes a string with {{key}}
func renderTemplate(t string, data map[string]interface{}) (string, error) {
	matches := templateVarRe.FindAllStringSubmatch(t, -1)

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
