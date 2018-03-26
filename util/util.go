package util

import (
	"strings"
	"time"
)

var monthMap map[string]int

func init() {
	monthMap = make(map[string]int, 1)
	monthMap["jan"] = 1
	monthMap["january"] = 1

	monthMap["feb"] = 2
	monthMap["february"] = 2

	monthMap["mar"] = 3
	monthMap["march"] = 3

	monthMap["apr"] = 4
	monthMap["april"] = 4

	monthMap["may"] = 5

	monthMap["jun"] = 6
	monthMap["june"] = 6

	monthMap["jul"] = 7
	monthMap["july"] = 7

	monthMap["aug"] = 8
	monthMap["august"] = 8

	monthMap["sep"] = 9
	monthMap["sept"] = 9
	monthMap["september"] = 9

	monthMap["oct"] = 10
	monthMap["october"] = 10

	monthMap["nov"] = 11
	monthMap["november"] = 11

	monthMap["dec"] = 12
	monthMap["december"] = 12
}

func GetTime(y, M, d, h, m, s int) int64 {
	return time.Date(y, time.Month(M), d, h, m, s, 0, time.UTC).Unix()
}

func GetMonthFromEn(month string) int {
	if _, ok := monthMap[strings.ToLower(month)]; ok {
		return monthMap[strings.ToLower(month)]
	}
	return -1
}
