package stream

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"

	"github.com/tidwall/gjson"
)

func NewFilterCommand() *cli.Command {
	return &cli.Command{
		Name:    "filter",
		Aliases: []string{"f"},
		Usage:   "使用JSONPath表达式过滤JSON对象",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "path", Aliases: []string{"p"}, Required: true, Usage: "JSONPath表达式"},
			&cli.StringFlag{Name: "input", Aliases: []string{"i"}, Required: true, Usage: "输入文件"},
			&cli.StringFlag{Name: "output", Aliases: []string{"o"}, Required: true, Usage: "输出文件"},
		},
		Action: func(c *cli.Context) error {
			inputFile := c.String("input")
			outputFile := c.String("output")
			jsonPath := c.String("path")

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
					continue
				}

				jsonValue, _ := json.Marshal(result.Value())
				_, err := writer.WriteString(string(jsonValue) + "\n")
				if err != nil {
					fmt.Fprintf(os.Stderr, "写入文件错误: %v\n", err)
					continue
				}
			}

			return scanner.Err()
		},
	}
}
