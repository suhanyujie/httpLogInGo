package config

import (
	"encoding/json"
	"flag"
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

var configPath *string
var ConfigData = &Config{}

func init()  {
	configPath = flag.String("c","../config.json","set config path:-c=config.json")
	flag.Parse()
	//解析配置文件
	ParseConfig(*configPath)
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
