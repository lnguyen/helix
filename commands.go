package main

import (
	"os"

	"github.com/longnguyen11288/helix/command"
	"github.com/longnguyen11288/helix/command/player"
	"github.com/mitchellh/cli"
)

// Commands is the mapping of all the available finger commands.
var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout}

	Commands = map[string]cli.CommandFactory{
		"player": func() (cli.Command, error) {
			return &player.Command{
				Ui: ui,
			}, nil
		},
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Version: Version,
				Ui:      ui,
			}, nil
		},
	}
}
