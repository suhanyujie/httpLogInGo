package parsers

import (
	"github.com/json-iterator/go"
	"log"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type LogInfo struct {
	Level       string `json:"level"`
	Time        string `json:"time"`
	Linenum     string `json:"linenum"`
	Msg         string `json:"msg"`
	ServiceName string `json:"serviceName"`
}

type LogMsg struct {
	Event    string `json:"event"`
	Key      string `json:"key"`
	Request  string `json:"request"`
	Response string `json:"response"`
	Type     string `json:"type"`
	Ip       string `json:"ip"`
	Time     string `json:"time"`
}

func ParseLogInfo(logContent []byte) (LogInfo,error) {
	var decodeData = LogInfo{}
	err := json.Unmarshal(logContent, &decodeData)
	if err!= nil {
		return decodeData,err
	}

	return decodeData,nil
}

func ParseLogMsg(msgJson string) *LogMsg {
	var logMsg1 = new(LogMsg)
	err := json.UnmarshalFromString(msgJson, logMsg1)
	if err != nil {
		log.Println(err)
	}

	return logMsg1
}

// 将字符串时间转换为时间格式类型
func ParseTime(timeStr string) (time.Time,error) {
	var (
		t1 time.Time
		err error
	)
	t1,err = time.Parse("2006-01-02 15:04:05", timeStr)
	if err!= nil {
		return t1,err
	}

	return t1.UTC(),err
}