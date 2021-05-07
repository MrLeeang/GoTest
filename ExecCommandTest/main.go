package main

import (
	"fmt"
	"os/exec"
)

func execCMD(name string, arg ...string) (bool, string) {

	var code = true

	cmd := exec.Command(name, arg...)

	ret, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		code = false
	}

	return code, string(ret)
}

func main() {
	code, ret := execCMD("sc.exe", "query", "RemoteAccess")
	if code != true {
		fmt.Println(ret)
		return
	}
	fmt.Println(ret)
	// s := string([]rune(ret))

	// countSplit := strings.Split(s, "\n")
	// fmt.Println(countSplit)
}
