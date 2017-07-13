package utils

import (
	"time"
)

//获取当前毫秒值
func GetCurrentMilliseconds() int64 {
	return time.Now().UnixNano() / 1000000
}

//获取当前秒数
func GetCurrentSeconds() int64 {
	return time.Now().Unix()
}

// sleep
func Sleep(waittime int) {
	time.Sleep(time.Duration(waittime) * time.Second)
}

// time
func Second(times int) time.Duration {
	return time.Duration(times) * time.Second
}

// get secord times
// 172606056
func GetSecordTimes() int64 {
	return time.Now().Unix()
}

//201611112113
func GetSecord2DateTimes(secord int64) string {
	tm := time.Unix(secord, 0)
	return tm.Format("20060102150405")

}

func GetDateTimes2Secord(datestring string) int64 {
	tm2, _ := time.Parse("20060102150405", datestring)
	return tm2.Unix()

}

func StringToMilliseconds(fmtstring string, datestring string) int64 {
	tm2, _ := time.Parse(fmtstring, datestring)
	return tm2.Unix()
}

func TodayString(level int) string {
	formats := "20060102150405"
	switch level {
	case 1:
		formats = "2006"
	case 2:
		formats = "200601"
	case 3:
		formats = "20060102"
	case 4:
		formats = "2006010215"
	case 5:
		formats = "200601021504"
	default:

	}
	return time.Now().Format(formats)
}
