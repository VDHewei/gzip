package stream

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"os"
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
			writer := bufio.NewWriter(writeFile)
			defer writer.Flush()
			// 读取所有行json字符串,并压缩
			result := gjson.GetBytes(jsonBytes, jsonPath)
			if !result.Exists() {
				log.Printf("查找目标不存在,json_path: %s\n", jsonPath)
				return nil
			}
			jsonValue, _ := json.Marshal(result.Value())
			if _, err = writer.WriteString(string(jsonValue) + "\n"); err != nil {
				log.Fatalf("写入文件错误: %v\n", err)
				return err
			}

			return nil
		},
	}
}
