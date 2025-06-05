package stream

import (
	"fmt"
	"github.com/chzyer/readline"
	"github.com/urfave/cli/v2"
	"log"
	"strings"
)

func NewInteractiveCommand() *cli.Command {
	return &cli.Command{
		Name:    "interactive",
		Aliases: []string{"i"},
		Usage:   "启动实时交互模式处理JSON数据",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "input", Aliases: []string{"i"}, Required: false, Usage: "输入文件"},
		},
		Action: func(c *cli.Context) error {
			inputFile := c.String("input")

			rl, err := readline.New("json-stream> ")
			if err != nil {
				return fmt.Errorf("无法初始化交互模式: %w", err)
			}
			defer rl.Close()

			for {
				cmd, err := rl.Readline()
				if err != nil { // io.EOF or canceled
					break
				}

				parts := strings.SplitN(cmd, " ", 2)
				if len(parts) < 2 {
					if cmd == "exit" || cmd == "q" {
						return nil
					}
					fmt.Println("无效命令格式。用法: [filter|replace|stats] [参数]")
					continue
				}

				switch parts[0] {
				case "filter":
					err = processFilterCommand(inputFile, parts[1])
				case "replace":
					err = processReplaceCommand(inputFile, parts[1])
				case "stats":
					err = processStatsCommand(inputFile, parts[1])
				case "exit", "q":
					return nil
				default:
					log.Println("未知命令:", parts[0])
					continue
				}
				if err != nil {
					log.Fatalf("执行错误: %v\n", err)
				}
			}

			return nil
		},
	}
}

func processFilterCommand(inputFile, args string) error {
	// 实现过滤命令的交互式处理
	log.Printf(`[Filter] inputFile= %s, args= %s\n`, inputFile, args)
	return nil
}

func processReplaceCommand(inputFile, args string) error {
	// 实现替换命令的交互式处理
	log.Printf(`[Replace] inputFile= %s, args= %s\n`, inputFile, args)
	return nil
}

func processStatsCommand(inputFile, args string) error {
	// 实现统计命令的交互式处理
	log.Printf(`[Stats] inputFile= %s, args= %s\n`, inputFile, args)
	return nil
}
