package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("date", "-s", "2021-05-18 13:55:00")
	ret, err := cmd.Output()
	fmt.Println(string(ret))
	fmt.Println(err)
}
