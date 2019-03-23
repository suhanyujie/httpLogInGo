package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	BaseUrl  string `json:"baseurl"`
	ApiToken string `json:"apitoken"`
	DB       DB     `json:"DB"`
}

type DB struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Passwd   string `json:"passwd"`
	Database string `json:"database"`
}

var ConfigData = &Config{}

func init()  {
	var (
		err error
	)
	_,err = ParseConfig("../config.json")
	if err!= nil {
		log.Println(err)
	}
}

// 解析json中的配置项
func ParseConfig(configPath string) (Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalln(err)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}
	data := &Config{}
	err = json.Unmarshal(content, ConfigData)
	if err != nil {
		log.Fatalln(err)
	}

	return *data, nil
}

func Get(key string, defaultValue interface{}) interface{}  {
	return ConfigData
}

func GetConfigData() *Config {
	return ConfigData
}
