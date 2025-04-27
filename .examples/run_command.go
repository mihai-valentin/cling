package example

import (
	"github.com/mihai-valentin/cling"
	"log"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatal("Missing required argument #1: command name")
	}

	// create the registry with a pre-filled commands list
	commandsRegistry := cling.NewRegistry(
		NewMyCommandWithArgsAndFlags(),
	)

	// or add command later
	commandsRegistry.Register(NewMyCommandWithoutArgsAndFlags())

	if err := commandsRegistry.RunCommand(args); err != nil {
		log.Fatal(err)
	}
}
