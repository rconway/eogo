package records

import (
	"flag"

	"github.com/rconway/eogo/utils"
)

func RecordsConnectCmdlineHandler(args []string) error {
	flagSet := flag.NewFlagSet("connect", flag.ExitOnError)
	var serverUrl string
	flagSet.StringVar(&serverUrl, "server", "http://localhost", "Specify the Catalogue server URL")
	return flagSet.Parse(args)
}

func RecordsCmdlineHandler(args []string) error {
	commands := utils.CommandMap{
		"connect": utils.CommandDetails{RecordsConnectCmdlineHandler, "connect to the Catalogue server"},
	}
	return utils.HandleCmdline(args, &commands)
}
