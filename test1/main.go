package main

import "fmt"

func main() {
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := range b {
		fmt.Println(i)
	}
}
