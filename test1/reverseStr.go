package main

import "fmt"

func main() {
	a := "abcd"
	r := []rune(a)
	b := []rune{}
	for i := 0; i < len(r); i++ {
		b = append(b, r[len(r)-i-1])
	}
	fmt.Println(string(b))
}
