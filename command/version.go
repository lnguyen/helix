package command

import "github.com/mitchellh/cli"

// VersionCommand is a Command implementation prints the version.
type VersionCommand struct {
	Version string
	Ui      cli.Ui
}

//Help command for version
func (c *VersionCommand) Help() string {
	return ""
}

//Run to print out version
func (c *VersionCommand) Run(_ []string) int {
	c.Ui.Output("Helix v" + c.Version)

	return 0
}

//Synopsis of command
func (c *VersionCommand) Synopsis() string {
	return "Prints the Helix version"
}
