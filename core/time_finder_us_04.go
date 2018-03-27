package core

import (
	"regexp"
	"strconv"

	util "github.com/Anderson-Lu/go_time_finder/util"
)

type TimeFinderUs04 struct {
	regexs []*regexp.Regexp
	err    error
}

func (self *TimeFinderUs04) GetName() string {
	return "[US-04]"
}

func (self *TimeFinderUs04) init() {
	regexs := []*regexp.Regexp{
		regexp.MustCompile(`([\d]{1,2})-([\d]{1,2})-([\d]{4})`),
		regexp.MustCompile(`([\d]{1,2})/([\d]{1,2})/([\d]{4})`),
		regexp.MustCompile(`([\d]{1,2})\.([\d]{1,2})\.([\d]{4})`),
	}
	self.regexs = regexs
}

func NewTimeFindeUs04() *TimeFinderUs04 {
	self := &TimeFinderUs04{}
	self.init()
	return self
}

func (self *TimeFinderUs04) Try(source string) []FinderResult {
	ret := []FinderResult{}
	for _, value := range self.regexs {
		ts := value.FindAllStringSubmatch(source, -1)
		for _, value2 := range ts {
			if len(value2) != 4 {
				continue
			}
			year, _ := strconv.Atoi(value2[3])
			month, _ := strconv.Atoi(value2[2])
			day, _ := strconv.Atoi(value2[1])
			timeInt := util.GetTime(year, month, day, 0, 0, 0)
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
