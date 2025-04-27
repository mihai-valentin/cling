package example

import (
	"fmt"
	"github.com/mihai-valentin/cling"
	"github.com/mihai-valentin/cling/arg"
)

type MyCommandWithArgsAndFlags struct {
	*cling.Command
	Verbose *cling.Flag
	Arg1    *cling.Arg
	Arg2    *cling.Arg
	Arg3    *cling.Arg
	Arg4    *cling.Arg
}

func NewMyCommandWithArgsAndFlags() *MyCommandWithArgsAndFlags {
	return &MyCommandWithArgsAndFlags{
		Command: cling.NewCommand("my-command-with-args-and-flags",
			cling.WithDescription("Description"),
		),
		Verbose: cling.NewFlag("--verbose"),
		Arg1:    cling.NewArg("a", arg.Min(1)),
		Arg2:    cling.NewArg("b", arg.Max(100)),
		Arg3:    cling.NewArg("c", arg.Range(10, 15)),
		Arg4:    cling.NewArg("d", arg.Blacklist("foo", "bar")),
	}
}

func (c *MyCommandWithArgsAndFlags) GetArgs() []*cling.Arg {
	return []*cling.Arg{c.Arg1, c.Arg2, c.Arg3, c.Arg4}
}

func (c *MyCommandWithArgsAndFlags) GetFlags() []*cling.Flag {
	return []*cling.Flag{c.Verbose}
}

func (c *MyCommandWithArgsAndFlags) Execute(input *cling.Input) error {
	if err := cling.UnmarshalArgsAndFlags(input, c); err != nil {
		return err
	}

	out := cling.NewOutput(c.Name)

	out.Log("Command started")
	out.LogIf("Verbose mode enabled", c.Verbose.Enabled)

	out.Log(fmt.Sprintf("Arg1: %d", c.Arg1.ValueAsInt()))
	out.Log(fmt.Sprintf("Arg2: %d", c.Arg2.ValueAsInt()))
	out.Log(fmt.Sprintf("Arg3: %d", c.Arg3.ValueAsInt()))
	out.Log(fmt.Sprintf("Arg4: %s", c.Arg4.Value))

	return nil
}
