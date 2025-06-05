package stream

import (
	"bufio"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"github.com/urfave/cli/v2"
	"os"
)

func NewReplaceCommand() *cli.Command {
	return &cli.Command{
		Name:    "replace",
		Aliases: []string{"r"},
		Usage:   "替换匹配JSONPath表达式的值",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "path", Aliases: []string{"p"}, Required: true, Usage: "JSONPath表达式"},
			&cli.StringFlag{Name: "input", Aliases: []string{"i"}, Required: true, Usage: "输入文件"},
			&cli.StringFlag{Name: "output", Aliases: []string{"o"}, Required: true, Usage: "输出文件"},
			&cli.StringFlag{Name: "value", Aliases: []string{"v"}, Required: true, Usage: "替换的新值"},
		},
		Action: func(c *cli.Context) error {
			inputFile := c.String("input")
			outputFile := c.String("output")
			jsonPath := c.String("path")
			newValue := c.String("value")

			readFile, err := os.Open(inputFile)
			if err != nil {
				return fmt.Errorf("打开输入文件失败: %w", err)
			}
			defer readFile.Close()

			writeFile, err := os.Create(outputFile)
			if err != nil {
				return fmt.Errorf("创建输出文件失败: %w", err)
			}
			defer writeFile.Close()

			scanner := bufio.NewScanner(readFile)
			writer := bufio.NewWriter(writeFile)
			defer writer.Flush()

			for scanner.Scan() {
				line := scanner.Text()
				result := gjson.Get(line, jsonPath)
				if !result.Exists() {
					_, err := writer.WriteString(line + "\n")
					if err != nil {
						fmt.Fprintf(os.Stderr, "写入文件错误: %v\n", err)
					}
					continue
				}

				// 替换指定路径的值
				modified, _ := sjson.Set(line, jsonPath, newValue)
				_, err := writer.WriteString(modified + "\n")
				if err != nil {
					fmt.Fprintf(os.Stderr, "写入文件错误: %v\n", err)
					continue
				}
			}

			return scanner.Err()
		},
	}
}
