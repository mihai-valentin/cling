package arg

import (
	"fmt"
	"strconv"
)

type MinValidationRule struct {
	min int
}

func Min(min int) *MinValidationRule {
	return &MinValidationRule{min}
}

func (r *MinValidationRule) Accepts(value string) (bool, error) {
	valueAsInt, err := strconv.Atoi(value)
	if err != nil {
		return false, err
	}

	if valueAsInt < r.min {
		return false, fmt.Errorf("value [%s] must be greater or equeal to [%d]", value, r.min)
	}

	return true, nil
}

type MaxValidationRule struct {
	max int
}

func Max(max int) *MaxValidationRule {
	return &MaxValidationRule{max}
}

func (r *MaxValidationRule) Accepts(value string) (bool, error) {
	valueAsInt, err := strconv.Atoi(value)
	if err != nil {
		return false, err
	}

	if valueAsInt > r.max {
		return false, fmt.Errorf("value [%s] must be less or equeal to [%d]", value, r.max)
	}

	return true, nil
}

type RangeValidationRule struct {
	min *MinValidationRule
	max *MaxValidationRule
}

func Range(min int, max int) *RangeValidationRule {
	return &RangeValidationRule{
		Min(min),
		Max(max),
	}
}

func (r *RangeValidationRule) Accepts(value string) (bool, error) {
	if _, err := r.min.Accepts(value); err != nil {
		return false, err
	}

	if _, err := r.max.Accepts(value); err != nil {
		return false, err
	}

	return true, nil
}
