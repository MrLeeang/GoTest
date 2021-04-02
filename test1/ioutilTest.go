package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	/*iotuil包
	ReadFile()
	WriteFile()
	ReadDir()
	*/
	// 读取文件
	// filePath := "2.txt"
	// data, _ := ioutil.ReadFile(filePath)
	// fmt.Println(string(data))

	// s1 := "fdsafdas"
	// err := ioutil.WriteFile("3.txt", []byte(s1), os.ModePerm)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// 读字符串
	// r := strings.NewReader("fdsjfklsdajlfkjadsklfkldas")
	// data, _ := ioutil.ReadAll(r)
	// fmt.Println(string(data))

	// 读目录
	// fileInfos, _ := ioutil.ReadDir(".")
	// for i := 0; i < len(fileInfos); i++ {
	// 	fileInfo := fileInfos[i]
	// 	fmt.Println(fileInfo.Name(), fileInfo.IsDir())
	// }

	// 临时目录和文件
	// dir, _ := ioutil.TempDir(".", "Test")
	// fmt.Println(dir)
	// file, _ := ioutil.TempFile(".", "Test")
	// fmt.Println(file.Name())
	// defer os.Remove(dir)
	// defer os.Remove(file.Name())
	// defer file.Close()

	listFiles(`d:/`, 0)

}

// 目录遍历
func listFiles(dirName string, level int) {
	s := "|--"
	for i := 0; i < level; i++ {
		s = "|   " + s
	}
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println(err)
	}
	for _, fileInfo := range fileInfos {
		filePath := dirName + "/" + fileInfo.Name()
		fmt.Println(s, filePath)
		if fileInfo.IsDir() {
			listFiles(filePath, level+1)
		}
	}
}
