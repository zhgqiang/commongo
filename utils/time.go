package utils

import (
	"time"
)

// TimeYearUnix 获取年 1月1日 时间戳
func TimeYearUnix() int64 {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return 0
	}
	now := time.Now()
	return time.Date(now.Year(), 1, 1, 0, 0, 0, 0, loc).Unix()
}

// TimeYear 获取当前年份
func TimeYear() int {
	return time.Now().Year()
}

// TimeMonthUnix 获取当月 1日 时间戳
func TimeMonthUnix() int64 {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return 0
	}
	now := time.Now()
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, loc).Unix()
}

// TimeMonth 获取当前月份
func TimeMonth() time.Month {
	return time.Now().Month()
}

// TimeDayUnix 获取当天 0点 时间戳
func TimeDayUnix() int64 {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return 0
	}
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc).Unix()
}

// TimeDay 获取当前日期
func TimeDay() int {
	return time.Now().Day()
}

// TimeHourUnix 获取当前小时时间戳
func TimeHourUnix() int64 {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return 0
	}
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, loc).Unix()
}

// TimeHour 获取当前小时
func TimeHour() int {
	return time.Now().Hour()
}

// TimeUnixFormatString 将时间戳戳转为格式化时间戳
func TimeUnixFormatString(timestamp int64, f string) string {
	return time.Unix(timestamp, 0).Format(f)
}

// TimeUnixToMonth 时间戳所在月份
func TimeUnixToMonth(timestamp int64) time.Month {
	return time.Unix(timestamp, 0).Month()
}

// TimeUnixToDay 时间戳所在日
func TimeUnixToDay(timestamp int64) int {
	return time.Unix(timestamp, 0).Day()
}
