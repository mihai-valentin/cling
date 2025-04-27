package cling

import (
	"fmt"
	"strings"
)

type Input struct {
	Args  map[string]string
	Flags map[string]byte
}

func NewInput(rawInput []string) (*Input, error) {
	var err error

	input := &Input{}

	if len(rawInput) == 0 {
		return input, nil
	}

	input.Args, err = parseArgs(rawInput)
	if err != nil {
		return nil, fmt.Errorf("cannot parse args [%s]: %w", rawInput, err)
	}

	input.Flags = parseFlags(rawInput)

	return input, nil
}

func parseArgs(rawInput []string) (map[string]string, error) {
	args := make(map[string]string)

	for _, part := range rawInput {
		if !strings.Contains(part, "=") {
			continue
		}

		splitPart := strings.Split(part, "=")

		if len(splitPart) != 2 {
			return nil, fmt.Errorf("invalid argument: %s", rawInput)
		}

		if splitPart[0] == "" || splitPart[1] == "" {
			return nil, fmt.Errorf("invalid argument: %s", rawInput)
		}

		if strings.HasSuffix(splitPart[0], " ") || strings.HasPrefix(splitPart[1], " ") {
			return nil, fmt.Errorf("invalid argument: %s", rawInput)
		}

		argName := strings.TrimSpace(splitPart[0])
		argValue := strings.TrimSpace(splitPart[1])

		args[argName] = argValue
	}

	return args, nil
}

func parseFlags(rawInput []string) map[string]byte {
	flags := make(map[string]byte)

	for _, part := range rawInput {
		if len(part) < 3 || !strings.HasPrefix(part, "--") {
			continue
		}

		flag := strings.TrimSpace(part)
		flags[flag] = 1
	}

	return flags
}
