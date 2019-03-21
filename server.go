package main

import (
	"github.com/suhanyujie/telegramApi/apiImpl"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"httpLogInGo/parsers"
	"httpLogInGo/store"
	"log"
	"net/http"
	"os"
	"time"
)

var logger *zap.Logger
var logConChan chan string

// 记录一下运行的时间
var startTime = time.Now().Unix()

func main() {
	logConChan = make(chan string, 2048)
	logger = getLogger()
	go func(ch1 <-chan string) {
		var mdWrapStr = ""
		var logMsgJson = ""
		var logMsg = new(parsers.LogMsg);
		for {
			select {
			// 从接收到的通道中拿出日志
			case logMsgJson = <-ch1:
				// 1.记录到日志文件中
				logger.Info(logMsgJson)
				// 2.发送到telegram信息提醒群中
				// 解析内容，并将key拿出，作为md文档的标题
				logMsg = parsers.ParseLogMsg(logMsgJson)
				log.Println(logMsg.Key)
				mdWrapStr = "\n## "+logMsg.Key+" \n\n" +
					"```\n" +
					logMsgJson +
					"\n```\n"
				apiImpl.SendMessage(mdWrapStr)
				// 3.记录到数据库中
				_, err := store.InsertOneLogMsg(logMsgJson)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}(logConChan)
	//  启动http服务器
	http.HandleFunc("/log", HandleLog)
	log.Println("http server start in port 8002.")
	http.ListenAndServe(":8002", nil)
}

func HandleLog(w http.ResponseWriter, r *http.Request) {
	var resRes = `{"status":"1","msg":"ok!"}`
	var val = r.FormValue("value")
	log.Println(val)
	go func(val string) {
		logConChan <- val
	}(val)
	w.Write([]byte(resRes))
}

// 获取基于zap的日志器
func getLogger() *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   "./logFile/log1.log", // 日志文件路径
		MaxSize:    10,                   // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 20,                   // 日志文件最多保存多少个备份
		MaxAge:     2,                    // 文件最多保存多少天
		Compress:   true,                 // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:  "time",
		LevelKey: "level",
		NameKey:  "logger",
		//CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel,                                                                     // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构造日志
	logger = zap.New(core, caller, development, filed)

	return logger
}
