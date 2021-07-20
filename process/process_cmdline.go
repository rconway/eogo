package process

import (
	"flag"
	"io"
)

func ProcessCmdlineHandler(args []string, w io.Writer) error {
	flagSet := flag.NewFlagSet("process", flag.ExitOnError)
	var serverUrl string
	flagSet.StringVar(&serverUrl, "server", "http://localhost", "Specify the Process server URL")
	return flagSet.Parse(args)
}
