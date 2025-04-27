package cling

import (
	"errors"
	"fmt"
)

type CommandWithArgs interface {
	GetArgs() []*Arg
}

type CommandWithFlags interface {
	GetFlags() []*Flag
}

type CommandWithArgsAndFlags interface {
	CommandWithArgs
	CommandWithFlags
}

func UnmarshalArgs(input *Input, c CommandWithArgs) error {
	for _, a := range c.GetArgs() {
		argFromInput, ok := input.Args[a.Name]

		if !a.Optional && !ok {
			return errors.New("missing required argument: " + a.Name)
		}

		if _, err := a.Accepts(argFromInput); err != nil {
			return fmt.Errorf("invalid argument [%s] value: %w", a.Name, err)
		}

		if ok {
			a.Set(argFromInput)
		}
	}

	return nil
}

func UnmarshalFlags(input *Input, c CommandWithFlags) error {
	for _, f := range c.GetFlags() {
		if flagFromInput, ok := input.Flags[f.Name]; ok {
			f.Set(flagFromInput == 1)
		}
	}

	return nil
}

func UnmarshalArgsAndFlags(input *Input, c CommandWithArgsAndFlags) error {
	if err := UnmarshalArgs(input, c); err != nil {
		return err
	}

	if err := UnmarshalFlags(input, c); err != nil {
		return err
	}

	return nil
}
