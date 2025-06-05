package main

import (
	"github.com/VDHewei/gzip/pkg/stream"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var version = "v0.0.1"

func main() {
	app := &cli.App{
		Name:    "json-stream",
		Version: version,
		Usage:   "处理大型JSON文件的流式处理工具",
		Commands: []*cli.Command{
			stream.NewGzipCommand(),
			stream.NewFilterCommand(),
			stream.NewStatsCommand(),
			stream.NewReplaceCommand(),
			stream.NewInteractiveCommand(),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Failed to run app: %v\n", err)
	}
}
