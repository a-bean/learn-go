package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func findGoModRoot(startDir string) (string, error) {
	dir := startDir
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parentDir := filepath.Dir(dir)
		if parentDir == dir { // 已到达文件系统根目录
			break
		}
		dir = parentDir
	}
	return "", fmt.Errorf("go.mod not found")
}

func main() {
	cwd, _ := os.Getwd()
	root, err := findGoModRoot(cwd) // 查找包含 go.mod 的目录
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Project root (via go.mod):", root)
	}
}
