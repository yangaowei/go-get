package logs

import (
	//"log"
	//"os"
	"testing"
)

func TestLogs(t *testing.T) {
	Log.Debug(" *Start: %v", "baidu.com")
	Log.Debug(" *Start: %v", "baidu.com")
	Log.Emergency("emergency")
	Log.Alert("alert")
	Log.Critical("critical")
	Log.Error("error")
	Log.Warning("warning")
	Log.Notice("notice")
	Log.Informational("informational")
	Log.Debug("debug")
}
