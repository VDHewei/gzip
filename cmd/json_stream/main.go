package main

import (
	"fmt"
	"github.com/VDHewei/gzip/pkg/stream"
	"github.com/urfave/cli/v2"
	"os"
)

var version = "v0.0.1"

func main() {
	app := &cli.App{
		Name:    "json-stream",
		Version: version,
		Usage:   "处理大型JSON文件的流式处理工具",
		Commands: []*cli.Command{
			stream.NewFilterCommand(),
			stream.NewStatsCommand(),
			stream.NewReplaceCommand(),
			stream.NewInteractiveCommand(),
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
