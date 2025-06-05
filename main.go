package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func compressFile(srcPath string, dstPath string) error {
	// 打开源文件
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// 创建目标.gz文件
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// 创建gzip写入器
	gzWriter := gzip.NewWriter(dstFile)
	defer gzWriter.Close()

	// 设置gzip头部信息（可选）
	gzWriter.Name = filepath.Base(srcPath)
	gzWriter.Comment = "Compressed by Go gzip package"
	gzWriter.ModTime = time.Now()

	// 执行压缩
	_, err = io.Copy(gzWriter, srcFile)
	return err
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("用法: gzip.exe <源文件> <目标文件.gz>")
	}

	err := compressFile(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal("压缩失败:", err)
	}
	log.Println("文件压缩成功")
}
