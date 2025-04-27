package arg

import "testing"

func TestBlacklistValidationRule_Accepts(t *testing.T) {
	t.Run("empty blacklist", func(t *testing.T) {
		rule := Blacklist()
		if ok, err := rule.Accepts("value"); !ok || err != nil {
			t.Errorf("expected rule to accept value, got %t, %s", ok, err)
		}
	})

	t.Run("non-empty blacklist", func(t *testing.T) {
		rule := Blacklist("black")

		if ok, err := rule.Accepts("value"); !ok || err != nil {
			t.Errorf("expected rule to accept value, got %t, %s", ok, err)
		}

		if ok, err := rule.Accepts("black"); ok || err == nil {
			t.Errorf("expected rule to not accept black, got %t, %s", ok, err)
		}
	})
}
