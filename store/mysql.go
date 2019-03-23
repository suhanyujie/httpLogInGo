package store

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"httpLogInGo/parsers"
	"log"
	"time"
)

var engine *xorm.Engine
var prefixStr = "log_"

// 初始化数据库
func init() {
	var err error
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", "root", "123456", "127.0.0.1", "3306", "logger")
	engine, err = xorm.NewEngine("mysql", dataSource)
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, prefixStr)
	engine.SetTableMapper(tbMapper)
	if err != nil {
		log.Fatalln(err)
	}
}

// 日志内容表1
type Content1 struct {
	Id          uint64    `xorm:"id pk not null autoincr "`                     // comment '主键'
	Event       string    `xorm:"'event' varchar(100) notnull "`                // comment '日志事件名称'
	Type        string    `xorm:"'type' varchar(18) notnull "`                  // comment '日志级别'
	Key         string    `xorm:"'key' index(key_index) varchar(200) notnull "` // comment '日志key，标签'，加索引
	Content     string    `xorm:"'content' text notnull "`                      // comment '日志内容'
	Request     string    `xorm:"'request' text notnull"`                       // comment '日志字段request'
	Response    string    `xorm:"'response' text notnull"`                      // comment '日志字段response'
	Ip          string    `xorm:"'ip' varchar(20) notnull default ''"`          //  comment '日志来源ip'
	Create_time time.Time `xorm:"create_time" `                                 // comment '日志新增时间'
	Update_time time.Time `xorm:"update_time" `                                 // comment '日志更新时间'
	DeleteFlag  int       `xorm:"delete_flag"`                                  //  comment '删除标记'
}

type LogContent Content1

func InsertOneLog(logContent string) (effectNum int64, err error) {
	log1 := parsers.LogInfo{}
	log1, err = parsers.ParseLogInfo([]byte(logContent))
	if err != nil {
		log.Println(err)
	}
	return InsertOneLogMsg(log1.Msg)
}

// 向数据就中插入一条数据
func InsertOneLogMsg(msgContent string) (effectNum int64, err error) {
	logMsg1 := new(parsers.LogMsg)
	// 将content转为结构体 LogMsg，便于获取其中的信息
	logMsg1 = parsers.ParseLogMsg(msgContent)
	insertData := new(Content1)
	if err != nil {
		return 0, err
	}
	// 字符串转为时间格式
	t1, err := parsers.ParseTime(logMsg1.Time)
	if err != nil {
		return 0, err;
	}
	insertData.Event = logMsg1.Event
	insertData.Type = logMsg1.Type
	insertData.Key = logMsg1.Key
	insertData.Content = msgContent
	insertData.Request = logMsg1.Request
	insertData.Response = logMsg1.Response
	insertData.Ip = logMsg1.Ip
	insertData.Create_time = t1
	insertData.Update_time = t1
	effectNum, err = engine.Insert(insertData)
	if err != nil {
		return 0, err;
	}

	return effectNum, err
}

// 如果表存在，则跳过，否则创建表
func CheckTableIsOk() (bool, error) {
	var err error
	has, err := engine.IsTableExist(prefixStr + "content1")
	if has == true {
		log.Println("table is exist,skip.")
		return true, err
	}
	err = engine.Sync2(new(Content1))
	if err != nil {
		log.Fatalln(err)
	}

	return true, err
}
