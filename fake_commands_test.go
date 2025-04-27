package cling

import "github.com/mihai-valentin/cling/arg"

type FakeCommand struct {
	*Command
}

func NewFakeCommand() *FakeCommand {
	return &FakeCommand{
		Command: NewCommand("fake-command"),
	}
}

func (c *FakeCommand) Execute(*Input) error {
	return nil
}

type FakeCommandWithOptionalArg struct {
	*Command
	Arg1 *Arg
}

func NewFakeCommandWithOptionalArg() *FakeCommandWithOptionalArg {
	return &FakeCommandWithOptionalArg{
		Command: NewCommand("fake-command"),
		Arg1:    NewOptionalArg("op_arg1"),
	}
}

func (c *FakeCommandWithOptionalArg) GetArgs() []*Arg {
	return []*Arg{c.Arg1}
}

func (c *FakeCommandWithOptionalArg) Execute(*Input) error {
	return nil
}

type FakeCommandWithArgsAndFlags struct {
	*Command
	Arg1  *Arg
	Arg2  *Arg
	Arg3  *Arg
	Flag1 *Flag
}

func NewFakeCommandWithArgsAndFlags() *FakeCommandWithArgsAndFlags {
	return &FakeCommandWithArgsAndFlags{
		Command: NewCommand("fake-command"),
		Arg1:    NewOptionalArg("op_arg1"),
		Arg2:    NewArg("arg2", arg.Blacklist("x", "z")),
		Arg3:    NewArg("arg3", arg.Range(0, 10)),
		Flag1:   NewFlag("flag1"),
	}
}

func (c *FakeCommandWithArgsAndFlags) GetArgs() []*Arg {
	return []*Arg{c.Arg1, c.Arg2, c.Arg3}
}

func (c *FakeCommandWithArgsAndFlags) GetFlags() []*Flag {
	return []*Flag{c.Flag1}
}

func (c *FakeCommandWithArgsAndFlags) Execute(*Input) error {
	return nil
}
