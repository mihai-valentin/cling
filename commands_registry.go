package cling

import (
	"fmt"
)

type NamedCommand interface {
	GetName() string
	Execute(input *Input) error
}

type NamedCommands map[string]NamedCommand

type CommandsRegistry struct {
	commands NamedCommands
}

func NewRegistry(commands ...NamedCommand) *CommandsRegistry {
	commandsMap := NamedCommands{}

	for _, command := range commands {
		commandsMap[command.GetName()] = command
	}

	return &CommandsRegistry{
		commands: commandsMap,
	}
}

func (r *CommandsRegistry) Register(c NamedCommand) *CommandsRegistry {
	r.commands[c.GetName()] = c

	return r
}

func (r *CommandsRegistry) RunCommand(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("too few arguments, expected at least 2, got %d", len(args))
	}

	commandName := args[1]

	c, err := r.getCommand(commandName)
	if err != nil {
		return fmt.Errorf("command [%s] not found: %s", commandName, err)
	}

	commandArgs := make([]string, 0)
	if len(args) > 2 {
		commandArgs = args[2:]
	}

	input, err := NewInput(commandArgs)
	if err != nil {
		return fmt.Errorf("cannot parse input: %s", err)
	}

	return c.Execute(input)
}

func (r *CommandsRegistry) getCommand(command string) (NamedCommand, error) {
	c, ok := r.commands[command]
	if !ok {
		return nil, fmt.Errorf("unknown command: %s", command)
	}

	return c, nil
}
