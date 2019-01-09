package cmd

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"
	"github.com/victorgama/slack-status/parser"
	"github.com/victorgama/slack-status/setter"
)

// UnsetStatus is responsible for unsetting the status on all slack teams
var UnsetStatus = cli.Command{
	Name:    "unset",
	Aliases: []string{"u"},
	Usage:   "Clears your Slack status",
	Action: func(c *cli.Context) error {
		config, err := parser.LoadConfig()
		if err != nil {
			return err
		}

		if err := setter.SetStatus(&parser.Status{}, config); err != nil {
			return err
		}

		fmt.Printf("Cleared your status on %s.\n", strings.Join(config.PrettyTeamNames(), ", "))

		return nil
	},
}
