package test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	envFilePath, _ := filepath.Abs("env.txt")
	fmt.Println("初始化环境变量", envFilePath)
	f, err := os.Open(envFilePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		arr := strings.Split(line, "=")
		if len(arr) == 2 {
			key := strings.TrimSpace(arr[0])
			value := strings.TrimSpace(arr[1])
			os.Setenv(key, value)
			fmt.Printf("set %s=%s\n", key, os.Getenv(key))
		}
		if err == io.EOF {
			break
		}
	}
	fmt.Println()
	// 执行单元测试
	m.Run()
}
