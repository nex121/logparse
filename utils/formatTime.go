package utils

import "time"

func SecondToTime(sec int64) string {
	return time.Unix(sec, 0).Format("2006-01-02 15:04:05")
}
