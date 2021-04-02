package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	srcFilePath := "1.jpg"
	destFilePath := "2.jpg"

	srcFile, _ := os.OpenFile(srcFilePath, os.O_RDONLY, os.ModePerm)
	destFile, _ := os.OpenFile(destFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)

	defer srcFile.Close()
	defer destFile.Close()

	fileInfo, _ := os.Stat(destFilePath)

	startSize := fileInfo.Size()

	fmt.Println(startSize)

	srcFile.Seek(startSize, io.SeekStart)

	byteData := make([]byte, 1024)
	for {
		n, err := srcFile.Read(byteData)
		if err == io.EOF || n == 0 {
			fmt.Println("文件复制完毕....", n)
			break
		}
		destFile.Write(byteData)
	}

	// LastIndex求最后一个/的index
	srcFilePath = "/etc/passwd"
	fileName := srcFilePath[strings.LastIndex(srcFilePath, "/")+1:]

	fmt.Println(fileName)
}
