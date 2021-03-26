package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func openRemoteAccess() {

	fmt.Println("openRemoteAccess ..... ")
	fmt.Println()

	cmd := exec.Command("sc.exe", "config", "RemoteAccess", "start=", "demand")
	cmd.Stdout = os.Stdout
	cmd.Run()
	cmd = exec.Command("sc.exe", "start", "RemoteAccess")
	cmd.Stdout = os.Stdout
	cmd.Run()
	// cmd = exec.Command("sc.exe", "query", "RemoteAccess")
	// cmd.Stdout = os.Stdout
	// cmd.Run()

}

func closeRemoteAccess() {

	fmt.Println("closeRemoteAccess ..... ")
	fmt.Println()

	cmd := exec.Command("sc.exe", "stop", "RemoteAccess")
	cmd.Stdout = os.Stdout
	cmd.Run()
	cmd = exec.Command("sc.exe", "config", "RemoteAccess", "start=", "disabled")
	cmd.Stdout = os.Stdout
	cmd.Run()

}

func main() {
	// 创建一个计时器
	timeTemplate := "2006-01-02 15:04:05"

	startTimeStr := os.Args[1]
	endTimeStr := os.Args[2]

	startTime, err := time.ParseInLocation(timeTemplate, startTimeStr, time.Local)

	if err != nil {
		panic(err.Error())
	}

	endTime, err := time.ParseInLocation(timeTemplate, endTimeStr, time.Local)
	if err != nil && endTime.IsZero() {
		panic(err.Error())
	}

	now := time.Now()

	startTicker := startTime.Sub(now)

	endTicker := endTime.Sub(now)

	fmt.Printf("RemoteAccess Service %s 之后开始启动....\n\n", startTicker)
	fmt.Printf("RemoteAccess Service %s 之后开始关闭....\n\n", endTicker)

	go func() {
		timeTicker := time.NewTicker(startTicker)
		for {
			<-timeTicker.C
			openRemoteAccess()
			break
		}
		timeTicker.Stop()
	}()

	go func() {
		timeTicker := time.NewTicker(endTicker)
		for {
			<-timeTicker.C
			closeRemoteAccess()
			break
		}
		timeTicker.Stop()
		os.Exit(0)
	}()

	select {}
}
