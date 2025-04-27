package cling

import (
	"errors"
	"testing"
)

type MockValidationRule struct {
	accepts bool
}

func (r *MockValidationRule) Accepts(string) (bool, error) {
	if !r.accepts {
		return false, errors.New("invalid value")
	}

	return true, nil
}

func TestArg_Accepts(t *testing.T) {
	t.Run("Arg without validation rules", func(t *testing.T) {
		arg := NewArg("test")

		if ok, err := arg.Accepts("value"); !ok || err != nil {
			t.Errorf("expected arg to accept value, got %t, %s", ok, err)
		}
	})

	t.Run("Arg with \"pass\" validation rules", func(t *testing.T) {
		arg := NewArg("test", &MockValidationRule{accepts: true})

		if ok, err := arg.Accepts("value"); !ok || err != nil {
			t.Errorf("expected arg to accept value, got %t, %s", ok, err)
		}
	})

	t.Run("Arg with \"fail\" validation rules", func(t *testing.T) {
		arg := NewArg("test", &MockValidationRule{accepts: false})

		if ok, err := arg.Accepts("value"); ok || err == nil {
			t.Errorf("expected arg to not accept value, got %t, %s", ok, err)
		}
	})

	t.Run("Arg with \"pass\" and \"fail\" validation rules", func(t *testing.T) {
		arg := NewArg(
			"test",
			&MockValidationRule{accepts: true},
			&MockValidationRule{accepts: false},
		)

		if ok, err := arg.Accepts("value"); ok || err == nil {
			t.Errorf("expected arg to not accept value, got %t, %s", ok, err)
		}
	})
}

func TestArg_ValueAsInt(t *testing.T) {
	t.Run("Arg with valid value", func(t *testing.T) {
		arg := NewArg("test")
		arg.Set("10")

		if value := arg.ValueAsInt(); value != 10 {
			t.Errorf("expected arg value to be 10, got %d", value)
		}
	})

	t.Run("Arg with invalid value", func(t *testing.T) {
		arg := NewArg("test")
		arg.Set("invalid")

		if value := arg.ValueAsInt(); value != 0 {
			t.Errorf("expected arg value to be 0, got %d", value)
		}
	})
}

func TestArg_Set(t *testing.T) {
	arg := NewArg("test")

	arg.Set("foo")
	if arg.Value != "foo" {
		t.Errorf("expected arg value to be \"foo\", got %s", arg.Value)
	}

	arg.Set("bar")
	if arg.Value != "bar" {
		t.Errorf("expected arg value to be \"bar\", got %s", arg.Value)
	}
}

func TestArg_Required(t *testing.T) {
	t.Run("Arg is required", func(t *testing.T) {
		arg := NewArg("test")

		if arg.Optional != false {
			t.Error("expected arg to be required, got optional")
		}
	})
}

func TestArg_Optional(t *testing.T) {
	t.Run("Arg is optional", func(t *testing.T) {
		arg := NewOptionalArg("test")

		if arg.Optional != true {
			t.Error("expected arg to be optional, got required")
		}
	})
}
