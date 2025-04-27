# Cling

[![Build&Test](https://github.com/mihai-valentin/cling/actions/workflows/test.yml/badge.svg)](https://github.com/mihai-valentin/cling/actions/workflows/test.yml)

Cling is a tiny CLI framework.

Cling treats CLI commands as Golang structs, providing full control over arguments and flags.
Also, Cling includes a simple validation system and a "commands' registry" abstraction.
Commands' registry simplifies the process of running commands using `os.Args`.

## Install

```shell
go get github.com/mihai-valentin/cling
```

## Convention

1. The command name must be passed as the second argument (e.g. `myapp command-name arg1=1 arg2=2`)
2. All command's args must be named (e.g. `arg_name=arg_value`)
3. All flags must be prefixed with a double dash (`--`)

## How to

1. Create a command struct based on the `cling.Command` abstraction
2. Create a constructor function and declare args and flags. Optionally, add validation rules for args.
   * Implement `GetArgs() []*cling.Arg` method if your command has args - return command's args
   * Implement `GetFlags() []*cling.Flags` method if your command has flags - return command's flags
3. Implement `Execute(input *Input) error` method - add the command logic there
4. Create an empty commands registry or use an existing one. Register your command.
5. Get the `os.Args` and run the command using the registry.

To learn more see `.exmaples` directory.

## Unmarshal input struct into command

```go
// Unmarshal arg only
err := cling.UnmarshalArgs(input, command)

// Unmarshal flags only
err := cling.UnmarshalFlags(input, command)

// Unmarshal args and flags
err := cling.UnmarshalArgsAndFlags(input, command)
```

## Available args validation rules

1. `int` arg value must be greater or equal to - `arg.Min(<min_value>)`
2. `int` arg value must be less or equal to - `arg.Max(<max_value>)`
3. `int` arg value must be between `a` and `b` - `arg.Range(<min_value>, <max_value>)`
4. Blacklist `string` arg values - `arg.Blacklist("master", "prod", "etc")`
