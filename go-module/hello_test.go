package hello

import "testing"

// name ...
func TestHello(t *testing.T) {
	want := "Hello, world."
	got := Hello()
	if got != want {
		t.Errorf("error")
	}
}
