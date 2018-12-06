package log

import "github.com/astaxie/beego/logs"

var Logger *logs.BeeLogger

func init() {
	Logger = logs.NewLogger()
	Logger.SetLogger(logs.AdapterConsole)
	logs.EnableFuncCallDepth(true)
}
