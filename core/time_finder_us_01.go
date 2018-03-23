package core

import (
	"strconv"
	"regexp"
	util "go_time_finder/util"
	"fmt"
)

type TimeFinderUs01 struct {
	regexs  []*regexp.Regexp
	err    error
}

func (self *TimeFinderUs01) GetName() string{
	return "[ZH-01]"
}

func (self *TimeFinderUs01) init() {
	regexs := []*regexp.Regexp {
		regexp.MustCompile(`^((31(?! (FEB|APR|JUN|SEP|NOV)))|((30|29)(?! FEB))|(29(?= FEB (((1[6-9]|[2-9]\d)(0[48]|[2468][048]|[13579][26])|((16|[2468][048]|[3579][26])00)))))|(0?[1-9])|1\d|2[0-8]) (JAN|FEB|MAR|MAY|APR|JUL|JUN|AUG|OCT|SEP|NOV|DEC) ((1[6-9]|[2-9]\d)\d{2})$`),		
	}
	self.regexs = regexs
}

func NewTimeFindeUs01() *TimeFinderUs01 {
	self := &TimeFinderUs01{}
	self.init()
	return self
}

func (self *TimeFinderUs01) Try(source string) []FinderResult{
	ret := []FinderResult{}
	for _,value := range self.regexs {
		ts:=value.FindAllStringSubmatch(source,-1)	
		fmt.Println(ts)	
		for _,value2 := range ts{
			if len(value2)!= 7 {
				continue
			}
			year,_:=strconv.Atoi(value2[1])
			month,_ := strconv.Atoi(value2[2])
			day,_:= strconv.Atoi(value2[3])
			hour,_:=strconv.Atoi(value2[4])
			min,_:=strconv.Atoi(value2[5])
			sec,_:=strconv.Atoi(value2[6])
			timeInt := util.GetTime(year,month,day,hour,min,sec)
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
