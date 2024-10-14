package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name
// checking for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Luchi"
	want := regexp.MustCompile(`\b` + name + `\b`)
	message, err := Hello("Luchi")
	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello("Luchi") = %q, %v, want match for %#q, nil`, message, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
