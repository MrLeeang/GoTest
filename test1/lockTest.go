package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var getIPlock sync.Mutex

func getIp(ip int, wg *sync.WaitGroup) {
	getIPlock.Lock()
	time.Sleep(3 * time.Second)
	fmt.Println("getIp", ip, time.Now().Format("2006年1月2日 15点4分5秒"))
	getIPlock.Unlock()
	wg.Done()
}

var getIp1lock sync.Mutex

func getIp1(ip int, wg *sync.WaitGroup) {
	getIp1lock.Lock()
	time.Sleep(3 * time.Second)
	fmt.Println("getIp1", ip, time.Now().Format("2006年1月2日 15点4分5秒"))
	getIp1lock.Unlock()
	wg.Done()
}

func main() {

	runtime.GOMAXPROCS(2)

	var wg = &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go getIp(i, wg)
		wg.Add(1)
	}

	for i := 0; i < 10; i++ {
		go getIp1(i, wg)
		wg.Add(1)
	}

	wg.Wait()

}
