package util

import "time"

func GetNowTime() string {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	return timeStr
}
