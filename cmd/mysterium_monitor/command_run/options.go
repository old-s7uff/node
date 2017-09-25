package command_run

import (
	"errors"
	"flag"
)

type CommandOptions struct {
	DirectoryRuntime string
	Node             string
	NodeFile         string
	ResultFile       string
}

func ParseArguments(args []string) (options CommandOptions, err error) {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.StringVar(
		&options.DirectoryRuntime,
		"runtime-dir",
		".",
		"Runtime directory for temp files (should be writable)",
	)
	flags.StringVar(
		&options.Node,
		"node",
		"",
		"Node to be checked",
	)
	flags.StringVar(
		&options.NodeFile,
		"node-file",
		"",
		"File with node list to be checked",
	)
	flags.StringVar(
		&options.ResultFile,
		"result-file",
		"",
		"File where CSV Output should be written",
	)

	err = flags.Parse(args[1:])
	if err != nil {
		return
	}

	if options.Node == "" && options.NodeFile == "" {
		err = errors.New("Provide which nodes to monitor!")
		return
	}

	if options.ResultFile == "" && options.NodeFile != "" {
		options.ResultFile = options.NodeFile + ".csv"
	}
	if options.ResultFile == "" {
		err = errors.New("Missing result file!")
		return
	}

	return options, err
}
