package main

import "fmt"

func b(n int) {
	defer fmt.Println(n)
	n += 100
}

func main() {
	b(1)
}
