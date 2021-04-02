package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// filePath := "1.txt"
	// file, err := os.Open(filePath)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// defer file.Close()
	// bReader := bufio.NewReader(file)

	// byteData := make([]byte, 1024)
	// for {
	// 	n, err := bReader.Read(byteData)
	// 	if err == io.EOF || n == 0 {
	// 		fmt.Println("文件读取完毕....")
	// 		return
	// 	}
	// 	fmt.Println(string(byteData), n, err)
	// }

	// for {
	// 	data, flag, err := bReader.ReadLine()
	// 	if err == io.EOF {
	// 		return
	// 	}
	// 	fmt.Println(string(data), flag, err)
	// }

	// for {
	// 	s1, err := bReader.ReadString('\n')
	// 	if err == io.EOF && s1 == "" {
	// 		break
	// 	}
	// 	fmt.Println(s1, err)
	// }

	// fmt读取键盘输入，如果又空格会有问题，所以要用bufio的包
	// for {
	// 	s2 := ""
	// 	fmt.Scanln(&s2)
	// 	fmt.Println(s2)
	// }

	// bufio.NewReader读取键盘输入
	// for {
	// 	b2 := bufio.NewReader(os.Stdin)
	// 	s2, _ := b2.ReadString('\n')
	// 	fmt.Println(s2)
	// }

	// bufio.NewScanner读取键盘输入
	// b2 := bufio.NewScanner(os.Stdin)
	// for b2.Scan() {
	// 	fmt.Println(b2.Text())
	// }

	writeFilePath := "2.txt"
	file, err := os.OpenFile(writeFilePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()
	// n, err := w.WriteString("\nHello World")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(n)
	for i := 1; i <= 100; i++ {
		w.WriteString("\nHello World")
		w.WriteString(strconv.Itoa(i))
	}

}
