package greetings

import (
	"regexp"
	"testing"
)

// check if Hello returns a string that contains the name
func TestHelloName(t *testing.T) {
	name := "Alice"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("%v") = %q, %v, want match for %#q, nil`, name, msg, err, want)
	}
}

// check if Hello returns an error if name is an empty string
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
