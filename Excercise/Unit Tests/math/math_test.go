package math

import (
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(6, 4)
	want := 10
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestSubstract(t *testing.T) {
	got := Substract(6, 4)
	want := 2
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
