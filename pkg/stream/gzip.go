package stream

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli/v2"
)

// NewGzipCommand 提供gzip压缩功能的子命令
func NewGzipCommand() *cli.Command {
	return &cli.Command{
		Name:    "gzip",
		Aliases: []string{"gz"},
		Usage:   "压缩文件为gzip格式",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "input", Aliases: []string{"i"}, Required: true, Usage: "输入文件路径"},
			&cli.StringFlag{Name: "output", Aliases: []string{"o"}, Required: false, Usage: "输出文件路径"},
			&cli.IntFlag{Name: "level", Aliases: []string{"l"}, Value: gzip.BestSpeed, Usage: "压缩级别 (1-9)"},
		},
		Action: func(c *cli.Context) error {
			inputFile := c.String("input")
			outputFile := c.String("output")
			level := c.Int("level")
			if outputFile == "" {
				outputFile = inputFile + ".gz"
			}
			// 验证压缩级别
			if level < gzip.HuffmanOnly || level > gzip.BestCompression {
				return fmt.Errorf("压缩级别必须在 %d-%d 之间", gzip.HuffmanOnly, gzip.BestCompression)
			}

			// 打开输入文件
			srcFile, err := os.Open(inputFile)
			if err != nil {
				return fmt.Errorf("无法打开输入文件: %w", err)
			}
			defer srcFile.Close()

			// 创建输出文件
			dstFile, err := os.Create(outputFile)
			if err != nil {
				return fmt.Errorf("无法创建输出文件: %w", err)
			}
			defer dstFile.Close()

			// 创建gzip写入器
			gw, err := gzip.NewWriterLevel(dstFile, level)
			if err != nil {
				return fmt.Errorf("无法创建gzip写入器: %w", err)
			}
			defer gw.Close()

			// 复制数据并压缩
			if _, err := io.Copy(gw, srcFile); err != nil {
				return fmt.Errorf("压缩过程中发生错误: %w", err)
			}

			fmt.Printf("成功将 %s 压缩为 %s\n", inputFile, outputFile)
			return nil
		},
	}
}
