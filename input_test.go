package cling

import (
	"log"
	"testing"
)

func TestNewInput(t *testing.T) {
	t.Run("empty input", func(t *testing.T) {
		input, err := NewInput([]string{})
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		if input.Args != nil {
			t.Errorf("expected args to be nil, got %v", input.Args)
		}

		if input.Flags != nil {
			t.Errorf("expected flags to be nil, got %v", input.Flags)
		}
	})

	t.Run("input with args", func(t *testing.T) {
		input, err := NewInput([]string{"foo=0", "bar=1"})
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		if input.Args == nil || len(input.Args) == 0 {
			t.Errorf("expected args to be not an empty map")
		}

		if input.Args["foo"] != "0" || input.Args["bar"] != "1" {
			t.Errorf("expected args to be foo=0 and bar=1, got %v", input.Args)
		}
	})

	t.Run("input with malformed args", func(t *testing.T) {
		malformedArgs := []string{"foo=", "bar =1", "buz= 1"}

		for _, arg := range malformedArgs {
			if i, err := NewInput([]string{arg}); err == nil {
				log.Println(i.Args)
				t.Errorf("expected error for arg %s, got nil", arg)
			}
		}
	})

	t.Run("input with flags", func(t *testing.T) {
		input, err := NewInput([]string{"--foo", "--bar"})
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		if input.Flags == nil || len(input.Flags) == 0 {
			t.Errorf("expected flags to be not an empty map")
		}

		if input.Flags["--foo"] != 1 || input.Flags["--bar"] != 1 {
			t.Errorf("expected flags --foo --bar, got %v", input.Flags)
		}
	})

	t.Run("input with args and flags", func(t *testing.T) {
		input, err := NewInput([]string{"foo=0", "--bar", "fuz=1", "--baz"})
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		if input.Args == nil || len(input.Args) == 0 {
			t.Errorf("expected args to be not an empty map")
		}

		if input.Flags == nil || len(input.Flags) == 0 {
			t.Errorf("expected flags to be not an empty map")
		}

		if input.Args["foo"] != "0" || input.Args["fuz"] != "1" {
			t.Errorf("expected args to be foo=0 and fuz=1, got %v", input.Args)
		}

		if input.Flags["--bar"] != 1 || input.Flags["--baz"] != 1 {
			t.Errorf("expected flags --bar --baz, got %v", input.Flags)
		}
	})
}
