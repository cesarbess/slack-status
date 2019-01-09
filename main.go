package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
	"github.com/victorgama/slack-status/cmd"
	"github.com/victorgama/slack-status/parser"
	"github.com/victorgama/slack-status/setter"
)

func main() {
	app := cli.NewApp()
	app.Name = "slack-status"
	app.Version = "0.1.0"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Victor Gama",
			Email: "hey@vito.io",
		},
	}
	app.HelpName = "slack-status"
	app.Usage = "Sets your Slack Status"
	app.UsageText = "slack-status [command]"
	app.Commands = []cli.Command{
		cmd.ListStatuses,
		cmd.UnsetStatus,
	}
	app.Action = func(c *cli.Context) error {
		config, err := parser.LoadConfig()
		if err != nil {
			return err
		}
		narg := c.NArg()
		var status *parser.Status
		switch narg {
		case 0:
			def, ok := config.Statuses["default"]
			if !ok {
				return parser.ErrNoDefault
			}
			status = def.PickRandom()
		case 1:
			statusName := c.Args().First()
			statuses, ok := config.Statuses[statusName]

			if !ok {
				keys := []string{}
				for k := range config.Statuses {
					keys = append(keys, k)
				}
				return fmt.Errorf("Invalid status key. Available keys: %s", strings.Join(keys, ", "))
			}

			status = statuses.PickRandom()
		default:
			return fmt.Errorf("Invalid usage")
		}

		if err := setter.SetStatus(status, config); err != nil {
			return err
		}

		fmt.Printf("Set on %s: %s\n", strings.Join(config.PrettyTeamNames(), ", "), status.String())

		return nil
	}
	app.Run(os.Args)
}
