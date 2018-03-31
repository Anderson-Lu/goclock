package core

import (
	"regexp"
	"strconv"
	"time"

	util "github.com/Anderson-Lu/go_time_finder/util"
)

type TimeFinderZh04 struct {
	regexs []*regexp.Regexp
	err    error
}

func (self *TimeFinderZh04) GetName() string {
	return "[ZH-04]"
}

func (self *TimeFinderZh04) init() {
	regexs := []*regexp.Regexp{
		regexp.MustCompile(`([\d]{1,2})-([\d]{1,2}) ([\d]{1,2}):([\d]{1,2})`),
		regexp.MustCompile(`([\d]{1,2})\.([\d]{1,2}) ([\d]{1,2}):([\d]{1,2})`),
		regexp.MustCompile(`([\d]{1,2})月([\d]{1,2})日 ([\d]{1,2}):([\d]{1,2})`),
	}
	self.regexs = regexs
}

func NewTimeFindeZh04() *TimeFinderZh04 {
	self := &TimeFinderZh04{}
	self.init()
	return self
}

func (self *TimeFinderZh04) Try(source string) []FinderResult {
	ret := []FinderResult{}
	for _, value := range self.regexs {
		ts := value.FindAllStringSubmatch(source, -1)
		for _, value2 := range ts {
			if len(value2) != 5 {
				continue
			}
			year := time.Now().Year()
			month, _ := strconv.Atoi(value2[1])
			day, _ := strconv.Atoi(value2[2])
			hour, _ := strconv.Atoi(value2[3])
			minute, _ := strconv.Atoi(value2[4])
			timeInt := util.GetTime(year, month, day, hour, minute, 0)
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
