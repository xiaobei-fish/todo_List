package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogPath    = "logs/"
	LogName    = "log"
	FileExt    = "log"
	TimeFormat = "20060102"
)

// 日志路径
func getLogFilePath() string {
	return fmt.Sprintf("%s", LogPath)
}

// 日志全路径 路径+文件名
func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogName, time.Now().Format(TimeFormat), FileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return handle
}

// 创建文件夹目录
func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
