package entrypoint

import core "go_time_finder/core"
import "fmt"

var finders []core.TimeFinder

func init() {
	RegistFinders(
		core.NewTimeFindeZh01(),
		core.NewTimeFindeZh02(),
		core.NewTimeFindeZh03(),
		core.NewTimeFindeUs01(),
	)
}

func RegistFinders(fs ...core.TimeFinder) {
	for _, v := range fs {
		finders = append(finders, v) 
	}
}

func FindTime(source string) []core.FinderResult {
	ret := []core.FinderResult{}
	for _, v := range finders {
		fmt.Println(v.GetName())
		ret=v.Try(source)
		if len(ret)>0 {
			break
		}	
	}
	return ret
}


