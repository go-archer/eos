package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func ProjectName(dir string) string {
	mod, err := os.Open(dir + "/go.mod")
	if err != nil {
		fmt.Println("go.mod不存在", err)
		return ""
	}
	defer mod.Close()
	name := ""
	_, err = fmt.Fscanf(mod, "module %s", &name)
	if err != nil {
		fmt.Println("读取go.mod失败", err)
		return ""
	}
	return name
}

func CreateFile(dir string, filename string) *os.File {
	filePath := filepath.Join(dir, filename)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Fatalf("创建目录失败: %s:%v", dir, err)
	}
	stat, _ := os.Stat(filePath)
	if stat != nil {
		return nil
	}
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("创建文件失败: %s:%v", filePath, err)
	}
	return file
}
