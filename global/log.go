package global

import (
	"fmt"
	"log"

	"github.com/astaxie/beego/logs"
)

var _log *logs.BeeLogger
var _boolInit bool = false

func init() {
	_log = logs.NewLogger()
	_log.SetLogger(logs.AdapterFile, `{"filename":"809.log","level":7,"maxlines":0,"maxsize":2097152,"daily":true,"maxdays":10}`)
	_boolInit = true
}

const (
	Analysis = "Analysis" //数据分析
	TCP      = "Tcp"      //tcp服务
	BUS      = "Bus"      //业务
	DB       = "DB"       //数据库操作
	Biz      = "Biz"      //分发
	Test     = "Test"     //测试

	SVR809   = "SERVER"
	SUB809   = "CLIENT"
	UPDATA   = "UPHandler"
	DOWNDATA = "DownHandler"
	Cache    = "CacheBase"
)

//FLH  Format Log Head  格式化日志内容头
func F(module, logName, log string) string {
	return fmt.Sprintf("%s|%s %v", module, logName, log)
}

//Debug log debug
func Debug(f string, v ...interface{}) {
	if !check(f, v) {
		return
	}
	_log.Debug(f, v...)
}

//Info  log info
func Info(f string, v ...interface{}) {
	if !check(f, v) {
		return
	}
	_log.Info(f, v...)
}

//Warning log warning
func Warning(f string, v ...interface{}) {
	if !check(f, v) {
		return
	}
	_log.Warning(f, v...)
}

//Error log error
func Error(f string, v ...interface{}) {
	if !check(f, v) {
		return
	}
	_log.Error(f, v...)
}

func check(f string, v ...interface{}) bool {
	if _log == nil || !_boolInit {
		log.Printf(f, v)
		return false
	}
	return true
}
