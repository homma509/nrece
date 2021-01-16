package main

import (
	"fmt"
	"os"

	"github.com/homma509/nrece/command"
	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("nrece", "1.0.0")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"server": func() (cli.Command, error) {
			return &command.ServerCommand{}, nil
		},
	}

	exitCode, err := c.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't start nrece: %s", err)
	}

	os.Exit(exitCode)
}
