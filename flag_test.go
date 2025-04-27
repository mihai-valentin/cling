package cling

import "testing"

func TestFlag_Set(t *testing.T) {
	f := NewFlag("test")

	f.Set(true)
	if !f.Enabled {
		t.Error("expected flag to be enabled")
	}

	f.Set(false)
	if f.Enabled {
		t.Error("expected flag to be disabled")
	}
}
