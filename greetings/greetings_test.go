package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Philipp"
	want := regexp.MustCompile(`\b` + name + `\b`)
	message, error := Hello("Philipp")
	if !want.MatchString(message) || error != nil {
		t.Fatalf(`Hello("Philipp") = %q, %v, want match for %#q, nil`, message, error, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	message, error := Hello("")
	if message != "" || error == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, message, error)
	}
}

func TestHelloNames(t *testing.T) {
	names := []string{"Philipp", "Minnos", "Emma"}
	messages, error := Hellos(names)
	if len(messages) != 3 || error != nil {
		t.Fatalf(`Hello("Philipp", "Minnos", "Emma") = %q, %v, want "Philipp", "Minnos", "Emma", error`, messages, error)
	}
}
