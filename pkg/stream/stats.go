package stream

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"log"
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

			var (
				jsonBytes []byte
				data      interface{}
			)
			jsonBytes, err = io.ReadAll(readFile)
			if err != nil {
				log.Fatalf("读取文件错误: %v\n", err)
				return err
			}
			if err = json.Unmarshal(jsonBytes, &data); err != nil {
				log.Fatalf("解析JSON错误: %v\n", err)
				return err
			}
			if jsonBytes, err = json.Marshal(data); err != nil {
				log.Fatalf("JSON压缩失败:%v\n", err)
				return err
			}
			totalCount := 0
			objectRegex := regexp.MustCompile(`\{.*\\}|\[.*\\]`)

			result := gjson.GetBytes(jsonBytes, jsonPath)
			if !result.Exists() {
				log.Printf("查找目标不存在,json_path: %s\n", jsonPath)
				return nil
			}

			// 提取数组内容中的对象或数组数量
			matches := objectRegex.FindAllString(result.String(), -1)
			totalCount += len(matches)
			fmt.Printf("匹配路径 '%s' 的数组元素总数: %d\n", jsonPath, totalCount)
			return nil
		},
	}
}
