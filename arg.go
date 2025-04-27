package cling

import (
	"log"
	"strconv"
)

type ValidationRule interface {
	Accepts(value string) (bool, error)
}

type Arg struct {
	Name            string
	Value           string
	Optional        bool
	validationRules []ValidationRule
}

func NewArg(name string, validationRules ...ValidationRule) *Arg {
	return &Arg{
		Name:            name,
		Optional:        false,
		validationRules: validationRules,
	}
}

func NewOptionalArg(name string, validationRules ...ValidationRule) *Arg {
	return &Arg{
		Name:            name,
		Optional:        true,
		validationRules: validationRules,
	}
}

func (a *Arg) Set(value string) {
	a.Value = value
}

func (a *Arg) Accepts(value string) (bool, error) {
	for _, rule := range a.validationRules {
		if ok, err := rule.Accepts(value); !ok || err != nil {
			return false, err
		}
	}

	return true, nil
}

func (a *Arg) ValueAsInt() int {
	intValue, err := strconv.Atoi(a.Value)
	if err != nil {
		log.Printf("cannot convert arg [%s] value [%s] to int: %s", a.Name, a.Value, err)
		return 0
	}

	return intValue
}
