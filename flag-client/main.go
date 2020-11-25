package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/cookiejar"
)

// GetFlagInfo 获取flag
func GetFlagInfo(mac string) map[string]interface{} {
	// 短信服务请求地址
	verifyURI := "http://192.168.0.8:8080/flag/generate"
	// 申明并创建一个cookie
	var gCurCookieJar *cookiejar.Jar
	gCurCookieJar = new(cookiejar.Jar)
	// 创建一个http client
	httpClient := &http.Client{
		CheckRedirect: nil,
		Jar:           gCurCookieJar,
	}

	postData := map[string]string{
		"mac": mac,
	}
	postJSONStr, _ := json.Marshal(postData)
	// 请求
	httpResp, err := httpClient.Post(
		verifyURI,
		"application/json",
		bytes.NewBuffer([]byte(postJSONStr)),
	)
	if err != nil {
		fmt.Println(err)
		return map[string]interface{}{}
	}
	// 关闭请求
	defer httpResp.Body.Close()
	// 获取返回值
	body, _ := ioutil.ReadAll(httpResp.Body)
	// json转成map
	ret := map[string]interface{}{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		fmt.Println(err)
		return map[string]interface{}{}
	}
	// 判断返回值
	if ret["status"] != 200 {
		fmt.Println(err)
		return map[string]interface{}{}
	}
	retData := ret["ret_data"]
	return retData.(map[string]interface{})
}

// GetMacAddress 获取mac地址
func GetMacAddress() string {

	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for _, inter := range interfaces {
		// mac := inter.HardwareAddr //获取本机MAC地址
		fmt.Println(inter.Name)
		fmt.Println(inter.Index)
		fmt.Println(inter.HardwareAddr)
	}
	fmt.Println("no eth0")
	return ""
}

func main() {
	mac := GetMacAddress()
	fmt.Println(mac)
}
