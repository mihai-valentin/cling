package cling

import "testing"

func TestCommand_GetName(t *testing.T) {
	c := NewCommand("test")

	if c.GetName() != "test" {
		t.Errorf("expected command name to be \"test\", got %s", c.GetName())
	}
}

func TestCommand_WithDescription(t *testing.T) {
	c := NewCommand("test", WithDescription("Description"))

	if c.Description != "Description" {
		t.Errorf("expected command description to be \"Description\", got %s", c.Description)
	}
}
