package cling

import (
	"testing"
)

func TestUnmarshalArgsAndFlags(t *testing.T) {
	t.Run("unmarshal empty input", func(t *testing.T) {
		input := &Input{}
		command := NewFakeCommandWithOptionalArg()

		if err := UnmarshalArgs(input, command); err != nil {
			t.Errorf("expected no error, got %s", err)
		}
	})

	t.Run("unmarshal valid input", func(t *testing.T) {
		input := &Input{
			Args: map[string]string{
				"arg2": "y",
				"arg3": "3",
			},
			Flags: map[string]byte{
				"flag1": 1,
			},
		}

		command := NewFakeCommandWithArgsAndFlags()

		if err := UnmarshalArgsAndFlags(input, command); err != nil {
			t.Errorf("expected no error, got %s", err)
		}
	})

	t.Run("unmarshal input without required args", func(t *testing.T) {
		input := &Input{
			Args: map[string]string{
				"arg1": "1",
				"arg3": "3",
			},
			Flags: map[string]byte{
				"flag1": 1,
			},
		}

		command := NewFakeCommandWithArgsAndFlags()

		if err := UnmarshalArgsAndFlags(input, command); err == nil {
			t.Errorf("expected error for missing required arg2, got no error")
		}
	})

	t.Run("unmarshal input with invalid args", func(t *testing.T) {
		input := &Input{
			Args: map[string]string{
				"arg2": "x",
				"arg3": "3",
			},
			Flags: map[string]byte{
				"flag1": 1,
			},
		}

		command := NewFakeCommandWithArgsAndFlags()

		if err := UnmarshalArgsAndFlags(input, command); err == nil {
			t.Errorf("expected error for invalid arg2, got no error")
		}
	})
}
