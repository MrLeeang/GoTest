package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func openRemoteAccessV2() {

	fmt.Println("openRemoteAccess ..... ")

	cmd := exec.Command("sc.exe", "config", "RemoteAccess", "start=", "demand")
	cmd.Stdout = os.Stdout
	cmd.Run()
	cmd = exec.Command("sc.exe", "start", "RemoteAccess")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println("SUCCESS!!")
}

func closeRemoteAccessV2() {

	fmt.Println("closeRemoteAccess ..... ")

	cmd := exec.Command("sc.exe", "stop", "RemoteAccess")
	cmd.Stdout = os.Stdout
	cmd.Run()
	cmd = exec.Command("sc.exe", "config", "RemoteAccess", "start=", "disabled")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println("Waiting for RemoteAccess service to exit ..... ")
	time.Sleep(5 * time.Second)

	fmt.Println("SUCCESS!!")
}

func operationRemoteAccessV2(operation string) {
	fmt.Printf("%s RemoteAccess ..... \n", operation)

	cmd := exec.Command("sc.exe", operation, "RemoteAccess")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	operation := os.Args[1]

	if operation == "start" {
		openRemoteAccessV2()
	} else if operation == "stop" {
		closeRemoteAccessV2()
	} else {
		operationRemoteAccessV2(operation)
	}
}
