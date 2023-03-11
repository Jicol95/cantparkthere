package cantparkthere

import (
	"testing"
)

// TestCantParkThere calls cantparkthere.CantParkThere checking
// for a valid return value.
func TestCantParkThere(t *testing.T) {
	want := "You can't park there"
	msg := CantParkThere()

	if want != msg {
		t.Fatalf(`Got = %q, Wwant match for %#q`, msg, want)
	}
}
