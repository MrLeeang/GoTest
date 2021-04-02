package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	t1 := time.Now()
	fmt.Println(t1)
	fmt.Println(t1.Weekday())
	fmt.Println(t1.Unix())

	// 字符串转，十进制，int64
	i2, _ := strconv.ParseInt("100", 10, 64)
	fmt.Println(i2)

	s2 := strconv.FormatInt(i2, 10)
	fmt.Println(s2)

	s3 := strconv.Itoa(int(i2))
	i3, _ := strconv.Atoi(s3)
	fmt.Println(s3, i3)

	fileInfo, err := os.Stat(`main.go`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.Mode())
	fileName1 := "main.go"
	fmt.Println(filepath.IsAbs(fileName1))
	// 获取绝对路径
	filePath, _ := filepath.Abs(fileName1)
	fmt.Println(filePath)
	// 获取父目录
	fmt.Println(path.Join(filePath, ".."))

	os.MkdirAll("123", os.ModePerm)

	// 文件
	os.Create("123.txt")
	//打开文件， open打开是只读的
	f, _ := os.Open("123.txt")
	f.Close()
	// 可读写
	f1, _ := os.OpenFile("123.txt", os.O_RDWR, os.ModePerm)
	f1.Close()
	// 删除文件活文件夹
	os.RemoveAll("123")
	os.RemoveAll("123.txt")

}
