package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello Justin"
	if got := Hello(); got != want {
		t.Errorf("whoops %q %q", got, want)
	}
}

func TestMain(t *testing.T) {
	want := 42
	if got := main(); got != want {
		t.Errorf("nope %d %d", got, want)
	}
}
