package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	s1, _ := ioutil.ReadFile("1.txt")
	fmt.Println(string(s1))
	var valueMap = make(map[string]int)
	s2 := strings.Split(string(s1), " ")
	for _, r := range s2 {
		valueMap[r]++
	}
	fmt.Println(valueMap)

	keys := []string{}

	for key, _ := range valueMap {
		keys = append(keys, key)
	}

	// 冒泡排序
	for i := 0; i < len(keys); i++ {
		for j := i + 1; j < len(keys); j++ {
			if valueMap[keys[i]] < valueMap[keys[j]] {
				keys[i], keys[j] = keys[j], keys[i]
			}
		}
	}

	// 错误排序，排序过程中需要动态维护keys，所以不能直接循环切片keys
	// for index1, key1 := range keys {

	// 	for index2, key2 := range keys[index1+1:] {
	// 		if valueMap[key1] > valueMap[key2] {
	// 			keys[index1], keys[index2] = key1, key2
	// 		}
	// 	}
	// }
	fmt.Println(keys)
}
