package process

import "flag"

func ProcessCmdlineHandler(args []string) error {
	flagSet := flag.NewFlagSet("process", flag.ExitOnError)
	var serverUrl string
	flagSet.StringVar(&serverUrl, "server", "http://localhost", "Specify the Process server URL")
	return flagSet.Parse(args)
}
