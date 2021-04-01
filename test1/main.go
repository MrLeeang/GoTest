package main

import "fmt"

func main() {
	switch x := 5; x {
	default:
		fmt.Println(x)
	case 6:
		x += 10
		fmt.Println(x)
		// witch中的每个case最后都有一个隐藏的break,fallthrough表示取消break，代码继续执行下一个case里面的代码(不匹配下一个case的值)
		fallthrough
	case 7:
		x += 20
		fmt.Println(x)
	case 8:
		x += 20
		fmt.Println(x)
	}
}
