package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/urfave/cli"
	"github.com/victorgama/slack-status/parser"
)

// ListStatuses lists all known and valid statuses in local ~/.slack-status file
var ListStatuses = cli.Command{
	Name:    "list",
	Aliases: []string{"l"},
	Usage:   "Shows all statuses defined in ~/.slack-status",
	Action: func(c *cli.Context) error {
		config, err := parser.LoadConfig()
		if err != nil {
			return err
		}
		if len(config.Statuses) < 1 {
			fmt.Println("No statuses defined :(")
			return nil
		}

		boldCyan := color.New(color.FgCyan, color.Bold).SprintfFunc()

		for k, v := range config.Statuses {
			if len(v) == 1 {
				fmt.Printf("%s: %s\n", boldCyan(k), v[0].String())
			} else {
				fmt.Println(boldCyan(k))
				for _, s := range v {
					fmt.Printf("  - %s\n", s.String())
				}
			}
			fmt.Println()
		}

		return nil
	},
}
