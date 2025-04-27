package arg

import (
	"fmt"
)

type BlacklistValidationRule struct {
	blacklist []string
}

func Blacklist(blacklist ...string) *BlacklistValidationRule {
	return &BlacklistValidationRule{blacklist}
}

func (r *BlacklistValidationRule) Accepts(value string) (bool, error) {
	for _, b := range r.blacklist {
		if b == value {
			return false, fmt.Errorf("value [%s] is not allowed", value)
		}
	}

	return true, nil
}
