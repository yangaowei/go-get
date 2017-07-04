package config

import (
//"strings"
)

// 来自配置文件的配置项。
var (
	LOG_CAP            int64  = 10000 // 日志缓存的容量
	LOG_LEVEL          int    = 8     // 全局日志打印级别（亦是日志文件输出级别）
	LOG_CONSOLE_LEVEL  int    = 8     // 日志在控制台的显示级别
	LOG_FEEDBACK_LEVEL int    = 7     // 客户端反馈至服务端的日志级别
	LOG_LINEINFO       bool   = true  // 日志是否打印行信息
	LOG_SAVE           bool   = true  // 是否保存所有日志到本地文件
	LOG                string = "logs/goget.log"
	LOG_ASYNC          bool   = false //日志是否异步输出
)

// func logLevel(l string) int {
// 	switch strings.ToLower(l) {
// 	case "app":
// 		return logs.LevelApp
// 	case "emergency":
// 		return logs.LevelEmergency
// 	case "alert":
// 		return logs.LevelAlert
// 	case "critical":
// 		return logs.LevelCritical
// 	case "error":
// 		return logs.LevelError
// 	case "warning":
// 		return logs.LevelWarning
// 	case "notice":
// 		return logs.LevelNotice
// 	case "informational":
// 		return logs.LevelInformational
// 	case "info":
// 		return logs.LevelInformational
// 	case "debug":
// 		return logs.LevelDebug
// 	}
// 	return -10
// }
