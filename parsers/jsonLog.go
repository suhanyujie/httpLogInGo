package parsers

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type LogInfo struct {
	Level string `json:"level"`
	Time string `json:"time"`
	Linenum string `json:"linenum"`
	Msg string `json:"msg"`
	ServiceName string `json:"serviceName"`
}

func ParseLogInfo(logContent []byte) (LogInfo,error) {
	var decodeData = LogInfo{}
	err := json.Unmarshal(logContent, &decodeData)
	if err!= nil {
		return decodeData,err
	}

	return decodeData,nil
}