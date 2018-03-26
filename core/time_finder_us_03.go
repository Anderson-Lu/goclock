package core

import (
	util "go_time_finder/util"
	"regexp"
	"strconv"
)

type TimeFinderUs03 struct {
	regexs []*regexp.Regexp
	err    error
}

func (self *TimeFinderUs03) GetName() string {
	return "[Us-03]"
}

func (self *TimeFinderUs03) init() {
	regexs := []*regexp.Regexp{
		regexp.MustCompile(`([\d]{1,2})-([\d]{1,2})-([\d]{4})\s+([\d]{1,2}):([\d]{1,2})`),
		regexp.MustCompile(`([\d]{1,2})/([\d]{4})/([\d]{1,2})\s+([\d]{1,2}):([\d]{1,2})`),
		regexp.MustCompile(`([\d]{1,2})\.([\d]{4})\.([\d]{1,2})\s+([\d]{1,2}):([\d]{1,2})`),
		regexp.MustCompile(`([\d]{1,2})日([\d]{1,2})月([\d]{4})年\s+([\d]{1,2}):([\d]{1,2})`),
		regexp.MustCompile(`([\d]{1,2})日([\d]{1,2})月([\d]{4})年\s+([\d]{1,2})时([\d]{1,2})分`),
		regexp.MustCompile(`([\d]{1,2})日([\d]{1,2})月([\d]{4})年\s+([\d]{1,2})時([\d]{1,2})分`),
	}
	self.regexs = regexs
}

func NewTimeFindeUs03() *TimeFinderUs03 {
	self := &TimeFinderUs03{}
	self.init()
	return self
}

func (self *TimeFinderUs03) Try(source string) []FinderResult {
	ret := []FinderResult{}
	for _, value := range self.regexs {
		ts := value.FindAllStringSubmatch(source, -1)
		for _, value2 := range ts {
			if len(value2) != 6 {
				continue
			}
			year, _ := strconv.Atoi(value2[3])
			month, _ := strconv.Atoi(value2[2])
			day, _ := strconv.Atoi(value2[1])
			hour, _ := strconv.Atoi(value2[4])
			min, _ := strconv.Atoi(value2[5])
			timeInt := util.GetTime(year, month, day, hour, min, 0)
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
