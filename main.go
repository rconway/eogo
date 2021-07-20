package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/rconway/eogo/process"
	"github.com/rconway/eogo/records"
	"github.com/rconway/eogo/utils"
)

func usage(commands *utils.CommandMap, w io.Writer) {
	fmt.Fprintf(w, "Usage:\n%v <command> <options>\n", filepath.Base(os.Args[0]))
	fmt.Fprintln(w, "  <command> is one of...")
	for k, v := range *commands {
		fmt.Fprintf(w, "    %v: %v\n", k, v.Description)
	}
	fmt.Fprintln(w, "  <options> varies for each command")
}

func main() {
	commands := utils.CommandMap{
		"records": utils.CommandDetails{records.RecordsCmdlineHandler, "tooling for OGC API Records"},
		"process": utils.CommandDetails{process.ProcessCmdlineHandler, "tooling for OGC API Processes"},
	}
	err := utils.HandleCmdline(os.Args, &commands)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		usage(&commands, os.Stdout)
	}
}
