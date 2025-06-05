package stream

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"regexp"

	"github.com/tidwall/gjson"
)

func NewStatsCommand() *cli.Command {
	return &cli.Command{
		Name:    "stats",
		Aliases: []string{"s"},
		Usage:   "统计JSON数组字段的子对象数量",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "path", Aliases: []string{"p"}, Required: true, Usage: "JSONPath表达式指向数组字段"},
			&cli.StringFlag{Name: "input", Aliases: []string{"i"}, Required: true, Usage: "输入文件"},
		},
		Action: func(c *cli.Context) error {
			inputFile := c.String("input")
			jsonPath := c.String("path")

			readFile, err := os.Open(inputFile)
			if err != nil {
				return fmt.Errorf("打开输入文件失败: %w", err)
			}
			defer readFile.Close()

			scanner := bufio.NewScanner(readFile)
			totalCount := 0
			objectRegex := regexp.MustCompile(`\{.*\}|\[.*\]`)

			for scanner.Scan() {
				line := scanner.Text()
				result := gjson.Get(line, jsonPath)
				if !result.Exists() {
					continue
				}

				// 提取数组内容中的对象或数组数量
				matches := objectRegex.FindAllString(result.String(), -1)
				totalCount += len(matches)
			}

			fmt.Printf("匹配路径 '%s' 的数组元素总数: %d\n", jsonPath, totalCount)
			return scanner.Err()
		},
	}
}
