package player

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/mitchellh/cli"
)

//Command is command to run player
type Command struct {
	Ui     cli.Ui
	args   []string
	player *Player
}

func (c *Command) readConfig() *Config {
	var cmdConfig Config
	var configFile string
	cmdFlags := flag.NewFlagSet("player", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }

	cmdFlags.StringVar(&cmdConfig.DataDir, "data-dir", "", "data directory")
	cmdFlags.StringVar(&cmdConfig.Name, "name", "", "name")
	cmdFlags.StringVar(&configFile, "config-file", "", "config-file")

	cmdFlags.IntVar(&cmdConfig.Port, "port", 0, "port")

	if err := cmdFlags.Parse(c.args); err != nil {
		return nil
	}

	config := DefaultConfig()
	if configFile != "" {
		fileConfig, err := ReadConfig(configFile)
		if err != nil {
			c.Ui.Error(err.Error())
			return nil
		}
		config = MergeConfig(config, fileConfig)
	}
	config = MergeConfig(config, &cmdConfig)

	return config
}

func (c *Command) setupPlayer(config *Config) error {
	c.Ui.Output("Starting player...")
	player, err := Create(config)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error starting player: %s", err))
		return err
	}
	c.player = player
	return nil
}

//Help is help menu
func (c *Command) Help() string {
	helpText := `
Usage: helix player [options]
  Starts a player server to play videos.
Options:
  -config-file=/config-file Path to config
  -data-dir=/data-dir       Path to where videos are stored
  -name=player              Name of the player used for tagging
  -port=8821                Port that server is runnning on
`
	return strings.TrimSpace(helpText)
}

//Run starts player server
func (c *Command) Run(args []string) int {
	c.args = args
	config := c.readConfig()
	if config == nil {
		return 1
	}
	log.Printf("[DEBUG] %v\n", config)
	if err := c.setupPlayer(config); err != nil {
		return 1
	}
	c.player.server.StartServer()
	return 0
}

//Synopsis is quick readme of player
func (c *Command) Synopsis() string {
	return "Start a player server"
}
