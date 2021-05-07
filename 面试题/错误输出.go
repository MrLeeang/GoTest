package main

import "fmt"

func main() {
	var b error

	if b == nil {
		fmt.Println(1)
	} else if b != nil {
		fmt.Println(2)
	}
}
