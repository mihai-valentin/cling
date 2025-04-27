package example

import (
	"github.com/mihai-valentin/cling"
)

type MyCommandWithoutArgsAndFlags struct {
	*cling.Command
}

func NewMyCommandWithoutArgsAndFlags() *MyCommandWithoutArgsAndFlags {
	return &MyCommandWithoutArgsAndFlags{
		Command: cling.NewCommand("my-command-without-args-and-flags",
			cling.WithDescription("Description"),
		),
	}
}

func (c *MyCommandWithoutArgsAndFlags) Execute(*cling.Input) error {
	out := cling.NewOutput(c.Name)

	out.Log("Command started")

	// Do something

	out.Log("Command finished")

	return nil
}
