package util

import (
	"time"
)

func GetTime(y,M,d,h,m,s int) int64{
	return time.Date(y,time.Month(M),d,h,m,s,0,time.UTC).Unix()
}