package arg

import "testing"

func TestMinValidationRule_Accepts(t *testing.T) {
	rule := Min(10)

	if ok, err := rule.Accepts("10"); !ok || err != nil {
		t.Errorf("expected rule to accept 10, got %t, %s", ok, err)
	}

	if ok, err := rule.Accepts("11"); !ok || err != nil {
		t.Errorf("expected rule to accept 11, got %t, %s", ok, err)
	}

	if ok, err := rule.Accepts("9"); ok || err == nil {
		t.Errorf("expected rule to not accept 9, got %t, %s", ok, err)
	}

	if ok, err := rule.Accepts("invalid"); ok || err == nil {
		t.Errorf("expected rule to not accept invalid, got %t, %s", ok, err)
	}
}

func TestMaxValidationRule_Accepts(t *testing.T) {
	rule := Max(10)

	if ok, err := rule.Accepts("10"); !ok || err != nil {
		t.Errorf("expected rule to accept 10, got %t, %s", ok, err)
	}

	if ok, err := rule.Accepts("9"); !ok || err != nil {
		t.Errorf("expected rule to accept 9, got %t, %s", ok, err)
	}

	if ok, err := rule.Accepts("11"); ok || err == nil {
		t.Errorf("expected rule to not accept 11, got %t, %s", ok, err)
	}

	if ok, err := rule.Accepts("invalid"); ok || err == nil {
		t.Errorf("expected rule to not accept invalid, got %t, %s", ok, err)
	}
}

func TestRangeValidationRule_Accepts(t *testing.T) {
	rule := Range(10, 20)

	if ok, err := rule.Accepts("10"); !ok || err != nil {
		t.Errorf("expected rule to accept 10, got %t, %s", ok, err)
	}

	if ok, err := rule.Accepts("11"); !ok || err != nil {
		t.Errorf("expected rule to accept 11, got %t, %s", ok, err)
	}

	if ok, err := rule.Accepts("20"); !ok || err != nil {
		t.Errorf("expected rule to accept 20, got %t, %s", ok, err)
	}

	if ok, err := rule.Accepts("9"); ok || err == nil {
		t.Errorf("expected rule to not accept 9, got %t, %s", ok, err)
	}

	if ok, err := rule.Accepts("21"); ok || err == nil {
		t.Errorf("expected rule to not accept 21, got %t, %s", ok, err)
	}

	if ok, err := rule.Accepts("invalid"); ok || err == nil {
		t.Errorf("expected rule to not accept invalid, got %t, %s", ok, err)
	}
}
