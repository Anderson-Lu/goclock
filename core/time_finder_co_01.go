package core

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type TimeFinderCo01 struct {
	regexs []*regexp.Regexp
	err    error
}

func (self *TimeFinderCo01) GetName() string {
	return "[Co-01]"
}

func (self *TimeFinderCo01) init() {
	regexs := []*regexp.Regexp{
		regexp.MustCompile(`([\d]{1,2})\s*(s|seconds|second|sec|秒)\s*(前|ago)`),
		regexp.MustCompile(`([\d]{1,2})\s*(m|minute|minutes|min|分钟|分)\s*(前|ago)`),
		regexp.MustCompile(`([\d]{1,2})\s*(h|hour|hours|小时|时)\s*(前|ago)`),
		regexp.MustCompile(`([\d]{1,2})\s*(d|day|days|天)\s*(前|ago)`),
	}
	self.regexs = regexs
}

func NewTimeFindeCo01() *TimeFinderCo01 {
	self := &TimeFinderCo01{}
	self.init()
	return self
}

func (self *TimeFinderCo01) Try(source string) []FinderResult {
	ret := []FinderResult{}
	for _, value := range self.regexs {
		ts := value.FindAllStringSubmatch(source, -1)
		fmt.Println(ts)
		for _, value2 := range ts {
			l := len(value2)
			if l != 4 {
				continue
			}

			now := time.Now()
			raw := value2[2]
			if raw == "s" || raw == "second" || raw == "seconds" || raw == "sec" || raw == "秒" {
				seconds, _ := strconv.Atoi(value2[1])
				now = now.Add(-time.Second * time.Duration(seconds))
			} else if raw == "m" || raw == "min" || raw == "minutes" || raw == "minute" || raw == "分" || raw == "分钟" {
				min, _ := strconv.Atoi(value2[1])
				now = now.Add(-time.Minute * time.Duration(min))
			} else if raw == "h" || raw == "hour" || raw == "hours" || raw == "小时" || raw == "时" {
				hour, _ := strconv.Atoi(value2[1])
				now = now.Add(-time.Hour * time.Duration(hour))
			} else if raw == "d" || raw == "day" || raw == "days" || raw == "天" {
				day, _ := strconv.Atoi(value2[1])
				now = now.Add(-time.Hour * 24 * time.Duration(day))
			} else {
				continue
			}

			timeInt := now.Unix()
			tmp := &FinderResult{
				SourceStr: value2[0],
				ResultStr: value2[0],
				ResultUTC: timeInt,
			}
			ret = append(ret, *tmp)
			break
		}
	}
	return ret
}
