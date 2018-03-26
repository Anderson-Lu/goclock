package core

import (
	"fmt"
	util "go_time_finder/util"
	"regexp"
	"strconv"
)

type TimeFinderUs01 struct {
	regexs []*regexp.Regexp
	err    error
}

func (self *TimeFinderUs01) GetName() string {
	return "[US-01]"
}

func (self *TimeFinderUs01) init() {
	regexs := []*regexp.Regexp{
		regexp.MustCompile(`(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Otc|Nov|Dec)\.\s+(\d{1,2})[\.|,]\s+(\d{4})`),
		regexp.MustCompile(`(January|February|March|April|May|June|July|Aguest|September|October|November|December)\s+(\d{1,2})[\.|,]\s+(\d{4})`),
		regexp.MustCompile(`(\d{1,2})\s+(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Otc|Nov|Dec)\.\s+(\d{4})`),
		regexp.MustCompile(`(\d{1,2})\s+(January|February|March|April|May|June|July|Aguest|September|October|November|December)[\.|,]\s+(\d{4})`),
		regexp.MustCompile(`(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Otc|Nov|Dec)\.\s+(\d{1,2})th[\.|,]\s+(\d{4})`),
		regexp.MustCompile(`(January|February|March|April|May|June|July|Aguest|September|October|November|December)\s+(\d{1,2})(th|st|nd|rd)[\.|,]\s+(\d{4})`),
		regexp.MustCompile(`(\d{1,2})(th|st|nd|rd)\s+(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Otc|Nov|Dec)\.\s+(\d{4})`),
		regexp.MustCompile(`(\d{1,2})(th|st|nd|rd)\s+(January|February|March|April|May|June|July|Aguest|September|October|November|December)[\.|,]\s+(\d{4})`),
	}
	self.regexs = regexs
}

func NewTimeFindeUs01() *TimeFinderUs01 {
	self := &TimeFinderUs01{}
	self.init()
	return self
}

func (self *TimeFinderUs01) Try(source string) []FinderResult {
	ret := []FinderResult{}
	for _, value := range self.regexs {
		ts := value.FindAllStringSubmatch(source, -1)
		fmt.Println(ts)
		for _, value2 := range ts {
			l := len(value2)
			if l < 4 {
				continue
			}
			s := make(map[int]string, 1)
			i := 0
			for _, v := range value2 {
				if v == "th" || v == "st" || v == "nd" || v == "rd" {
					continue
				}
				s[i] = v
				i++
			}
			year, _ := strconv.Atoi(s[3])
			month := util.GetMonthFromEn(s[1])
			day, _ := strconv.Atoi(s[2])
			if month == -1 {
				continue
			}
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
