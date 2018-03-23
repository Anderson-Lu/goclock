package core

import (
	"strconv"
	"regexp"
	util "go_time_finder/util"
)

type TimeFinderZh03 struct {
	regexs  []*regexp.Regexp
	err    error
}

func (self *TimeFinderZh03) GetName() string{
	return "[ZH-03]"
}

func (self *TimeFinderZh03) init() {
	regexs := []*regexp.Regexp {
		regexp.MustCompile(`([\d]{4})-([\d]{1,2})-([\d]{1,2})`),
		regexp.MustCompile(`([\d]{4})/([\d]{1,2})/([\d]{1,2})`),
		regexp.MustCompile(`([\d]{4})\.([\d]{1,2})\.([\d]{1,2})`),
		regexp.MustCompile(`([\d]{4})年([\d]{1,2})月([\d]{1,2})日`),
	}
	self.regexs = regexs
}

func NewTimeFindeZh03() *TimeFinderZh03 {
	self := &TimeFinderZh03{}
	self.init()
	return self
}

func (self *TimeFinderZh03) Try(source string) []FinderResult{
	ret := []FinderResult{}
	for _,value := range self.regexs {
		ts:=value.FindAllStringSubmatch(source,-1)		
		for _,value2 := range ts{
			if len(value2)!= 4 {
				continue
			}
			year,_:=strconv.Atoi(value2[1])
			month,_ := strconv.Atoi(value2[2])
			day,_:= strconv.Atoi(value2[3])
			timeInt := util.GetTime(year,month,day,0,0,0)
			tmp := &FinderResult{
				SourceStr:value2[0],
				ResultStr:value2[0],
				ResultUTC:timeInt,
			}
			ret = append(ret,*tmp)
			break
		}
	}
	return ret
}
