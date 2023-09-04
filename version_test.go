package goutil

import "testing"

func TestVersion(t *testing.T) {
	got := Version()
	if got != "v1.0.0" {
		t.Errorf("Version() = %s; want 1", got)
	}
}
