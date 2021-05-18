package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// MyFormatter MyFormatter
type MyFormatter struct{}

// Format Format
func (s *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006/01/02 15:04:05")
	// 日志格式
	msg := fmt.Sprintf("%s [%s] %s:%v	%s\n", timestamp, strings.ToUpper(entry.Level.String()), entry.Caller.File, entry.Caller.Line, entry.Message)
	return []byte(msg), nil
}

// NewLogger 定义日志
func NewLogger() *logrus.Logger {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}
	logFileName := now.Format("2006-01-02") + ".log"
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//实例化
	logger := logrus.New()

	//设置输出控制和文件
	writers := []io.Writer{
		src,
		// os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)

	logger.SetOutput(fileAndStdoutWriter)

	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	logger.Writer()

	// 显示行号
	logger.SetReportCaller(true)

	// 设置日志格式
	logger.SetFormatter(new(MyFormatter))

	return logger
}

// Logger 打印日志调用
var Logger = NewLogger()

type pluginInfo struct {
	CNIVersion_        string   `json:"cniVersion"`
	SupportedVersions_ []string `json:"supportedVersions,omitempty"`
}

func main() {

	Logger.Info("fdsfsdfdsfdsfg")

	command := os.Getenv("CNI_COMMAND")

	switch command {
	case "ADD":
		_ = `{
			"cniVersion": "1.0.0",
			"interfaces": [                                            (this key omitted by IPAM plugins)
				{
					"name": "<name>",
					"mac": "<MAC address>",                            (required if L2 addresses are meaningful)
					"sandbox": "<netns path or hypervisor identifier>" (required for container/hypervisor interfaces, empty/omitted for host interfaces)
				}
			],
			"ips": [
				{
					"version": "<4-or-6>",
					"address": "<ip-and-prefix-in-CIDR>",
					"gateway": "<ip-address-of-the-gateway>",          (optional)
					"interface": <numeric index into 'interfaces' list>
				},
				...
			],
			"routes": [                                                (optional)
				{
					"dst": "<ip-and-prefix-in-cidr>",
					"gw": "<ip-of-next-hop>"                           (optional)
				},
				...
			],
			"dns": {                                                   (optional)
			  "nameservers": <list-of-nameservers>                     (optional)
			  "domain": <name-of-local-domain>                         (optional)
			  "search": <list-of-additional-search-domains>            (optional)
			  "options": <list-of-options>                             (optional)
			}
		  }`

		fmt.Println(`{"cniVersion": "1.0.0"}`)
	case "CHECK":

	case "DEL":

	case "VERSION":
		version := &pluginInfo{}
		version.CNIVersion_ = "1.0.0"
		version.SupportedVersions_ = []string{"0.1.0", "0.2.0", "0.3.0", "0.3.1", "0.4.0", "1.0.0"}
		json.NewEncoder(os.Stdout).Encode(version)
	default:
		os.Exit(1)
	}

	for _, e := range os.Environ() {

		Logger.Info(e)

	}

}
