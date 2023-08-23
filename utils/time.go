package utils

import "time"

func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}
func OneHourInSeconds() int64 {
	return int64(3600)
}

func OneWeekFromNow() int64 {

	return time.Now().Unix() + int64(7*24*time.Hour.Seconds())
}
