package process

import (
	"flag"
	"fmt"
	"io"
	"path/filepath"

	"github.com/rconway/eogo/utils"
)

func usage(args []string, commands *utils.CommandMap, w io.Writer) {
	fmt.Fprintf(w, "Usage:\n%v <command> <options>\n", filepath.Base(args[0]))
	fmt.Fprintln(w, "  <command> is one of...")
	for k, v := range *commands {
		fmt.Fprintf(w, "    %v: %v\n", k, v.Description)
	}
	fmt.Fprintln(w, "  <options> varies for each command")
}

// Expects: args[0] = 'connect', args[1:] = flags
func processConnectCmdlineHandler(args []string, w io.Writer) (utils.CommandExecutor, error) {
	flagSet := flag.NewFlagSet("connect", flag.ExitOnError)
	var serverUrl string
	flagSet.StringVar(&serverUrl, "server", "http://localhost", "Specify the Processing server URL")
	err := flagSet.Parse(args[1:])
	var cmdExe ProcessConnect
	if err == nil {
		cmdExe = ProcessConnect{Server: serverUrl}
	}
	return &cmdExe, err
}

// Expects: args[0] = 'process', args[1] = subcommand
func ProcessCmdlineHandler(args []string, w io.Writer) (utils.CommandExecutor, error) {
	commands := utils.CommandMap{
		"connect": utils.CommandDetails{Handler: processConnectCmdlineHandler, Description: "connect to the Processing server"},
	}
	return utils.HandleCmdline(args, &commands, usage, w)
}
