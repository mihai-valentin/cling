package cling

import (
	"testing"
)

func TestCommandsRegistry_RunCommand(t *testing.T) {
	t.Run("too few arguments", func(t *testing.T) {
		if err := NewRegistry().RunCommand([]string{}); err == nil {
			t.Error("expected error for too few arguments")
		}

		if err := NewRegistry().RunCommand([]string{"app"}); err == nil {
			t.Error("expected error for too few arguments")
		}
	})

	t.Run("undefined command", func(t *testing.T) {
		if err := NewRegistry().RunCommand([]string{"app", "undefined-command"}); err == nil {
			t.Error("expected error for undefined command")
		}
	})

	t.Run("malformed input args", func(t *testing.T) {
		if err := NewRegistry().RunCommand([]string{"app", "fake-command", "arg1 =0", "arg2= 1"}); err == nil {
			t.Error("expected error for malformed input args")
		}
	})

	t.Run("successful command execution", func(t *testing.T) {
		registry := NewRegistry()
		registry.Register(NewFakeCommand())

		if err := registry.RunCommand([]string{"app", "fake-command"}); err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		registry = NewRegistry(
			NewFakeCommand(),
		)

		if err := registry.RunCommand([]string{"app", "fake-command"}); err != nil {
			t.Errorf("expected no error, got %s", err)
		}
	})
}
