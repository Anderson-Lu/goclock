package core

import (
	"strconv"
	"regexp"
	util "go_time_finder/util"
)

type TimeFinderZh02 struct {
	regexs  []*regexp.Regexp
	err    error
}

func (self *TimeFinderZh02) GetName() string{
	return "[ZH-02]"
}

func (self *TimeFinderZh02) init() {
	regexs := []*regexp.Regexp {
		regexp.MustCompile(`([\d]{4})-([\d]{1,2})-([\d]{1,2})\s+([\d]{1,2}):([\d]{1,2})`),
		regexp.MustCompile(`([\d]{4})/([\d]{1,2})/([\d]{1,2})\s+([\d]{1,2}):([\d]{1,2})`),
		regexp.MustCompile(`([\d]{4})\.([\d]{1,2})\.([\d]{1,2})\s+([\d]{1,2}):([\d]{1,2})`),
		regexp.MustCompile(`([\d]{4})年([\d]{1,2})月([\d]{1,2})日\s+([\d]{1,2}):([\d]{1,2})`),
		regexp.MustCompile(`([\d]{4})年([\d]{1,2})月([\d]{1,2})日\s+([\d]{1,2})时([\d]{1,2})分`),
		regexp.MustCompile(`([\d]{4})年([\d]{1,2})月([\d]{1,2})日\s+([\d]{1,2})時([\d]{1,2})分`),
	}
	self.regexs = regexs
}

func NewTimeFindeZh02() *TimeFinderZh02 {
	self := &TimeFinderZh02{}
	self.init()
	return self
}

func (self *TimeFinderZh02) Try(source string) []FinderResult{
	ret := []FinderResult{}
	for _,value := range self.regexs {
		ts:=value.FindAllStringSubmatch(source,-1)		
		for _,value2 := range ts{
			if len(value2)!= 6 {
				continue
			}
			year,_:=strconv.Atoi(value2[1])
			month,_ := strconv.Atoi(value2[2])
			day,_:= strconv.Atoi(value2[3])
			hour,_:=strconv.Atoi(value2[4])
			min,_:=strconv.Atoi(value2[5])			
			timeInt := util.GetTime(year,month,day,hour,min,0)
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
