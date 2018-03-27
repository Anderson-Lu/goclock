package core

import (
	"regexp"
	"strconv"

	util "github.com/Anderson-Lu/go_time_finder/util"
)

type TimeFinderUs02 struct {
	regexs []*regexp.Regexp
	err    error
}

func (self *TimeFinderUs02) GetName() string {
	return "[US-02]"
}

func (self *TimeFinderUs02) init() {
	regexs := []*regexp.Regexp{
		regexp.MustCompile(`([\d]{1,2})\s+([\d]{1,2})\s+([\d]{4})\s+([\d]{1,2}):([\d]{1,2}):([\d]{1,2})?`),
		regexp.MustCompile(`([\d]{1,2})\.([\d]{1,2})\.([\d]{4})\s+([\d]{1,2}):([\d]{1,2}):([\d]{1,2})?`),
		regexp.MustCompile(`([\d]{1,2})-([\d]{1,2})-([\d]{4})\s+([\d]{1,2}):([\d]{1,2}):([\d]{1,2})?`),
	}
	self.regexs = regexs
}

func NewTimeFindeUs02() *TimeFinderUs02 {
	self := &TimeFinderUs02{}
	self.init()
	return self
}

func (self *TimeFinderUs02) Try(source string) []FinderResult {
	ret := []FinderResult{}
	for _, value := range self.regexs {
		ts := value.FindAllStringSubmatch(source, -1)
		for _, value2 := range ts {
			if len(value2) != 7 {
				continue
			}
			year, _ := strconv.Atoi(value2[3])
			month, _ := strconv.Atoi(value2[2])
			day, _ := strconv.Atoi(value2[1])
			hour, _ := strconv.Atoi(value2[4])
			min, _ := strconv.Atoi(value2[5])
			sec, _ := strconv.Atoi(value2[6])
			timeInt := util.GetTime(year, month, day, hour, min, sec)
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
