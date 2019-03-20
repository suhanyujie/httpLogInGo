package store

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"log"
	"time"
)

var engine *xorm.Engine

func init()  {
	var err error
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", "root", "123456", "127.0.0.1", "3306", "logger")
	engine, err = xorm.NewEngine("mysql", dataSource)
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "log_")
	engine.SetTableMapper(tbMapper)
	if err!= nil {
		log.Fatalln(err)
	}
}

type Content1 struct {
	Id int `xorm:"id pk not null autoincr "`// comment '主键'
	Content string `xorm:"'content' text notnull "`// comment '日志内容'
	Ip string `xorm:"'ip' varchar(12) notnull default ''"`//  comment '日志来源ip'
	OrderId string `xorm:"'order_id' varchar(30) notnull default '' "`// comment '订单号'
	Create_time time.Time `xorm:"create_time" `// comment '日志新增时间'
	Update_time time.Time `xorm:"update_time" `// comment '日志更新时间'
	DeleteFlag int `xorm:"delete_flag"`//  comment '删除标记'
}

func CheckTableIsOk() (bool,error) {
	var err error
	err = engine.Sync2(new(Content1))
	if err!= nil {
		log.Fatalln(err)
	}

	return true,err
}
