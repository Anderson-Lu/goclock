package entrypoint

import (
	"reflect"
	"testing"

	core "github.com/Anderson-Lu/go_time_finder/core"
)

type a struct {
	source string
}

type demo struct {
	name string
	args a
	want []core.FinderResult
}

func buildDemo(format, args string, results ...int64) demo {
	want := []core.FinderResult{}
	for _, v := range results {
		item := core.FinderResult{
			SourceStr: args,
			ResultStr: args,
			ResultUTC: v,
		}
		want = append(want, item)
	}
	r := demo{
		name: "[TESTING] " + format,
		args: a{source: args},
		want: want,
	}
	return r
}

func TestFindTime(t *testing.T) {

	tests := []demo{
		//zh-01
		buildDemo("yyyy-MM-dd hh:mm:ss", "2018-01-01 12:12:12", 1514808732),
		buildDemo("yyyy/MM/dd hh:mm:ss", "2018/01/01 12:12:12", 1514808732),
		buildDemo("yyyy.MM.dd hh:mm:ss", "2018.01.01 12:12:12", 1514808732),
		buildDemo("yyyy年MM月dd日 hh:mm:ss", "2018年01月01日 12:12:12", 1514808732),
		buildDemo("yyyy年MM月dd日 hh时mm分ss秒", "2018年01月01日 12时12分12秒", 1514808732),
		buildDemo("yyyy年MM月dd日 hh時mm分ss秒", "2018年01月01日 12時12分12秒", 1514808732),
		//zh-02
		buildDemo("yyyy-MM-dd hh:mm", "2018-01-01 12:12", 1514808720),
		buildDemo("yyyy/MM/dd hh:mm", "2018/01/01 12:12", 1514808720),
		buildDemo("yyyy.MM.dd hh:mm", "2018.01.01 12:12", 1514808720),
		buildDemo("yyyy年MM月dd日 hh:mm", "2018年01月01日 12:12", 1514808720),
		buildDemo("yyyy年MM月dd日 hh时mm分", "2018年01月01日 12时12分", 1514808720),
		buildDemo("yyyy年MM月dd日 hh時mm分", "2018年01月01日 12時12分", 1514808720),
		//zh-03
		buildDemo("yyyy-MM-dd", "2018-01-01", 1514764800),
		buildDemo("yyyy/MM/dd", "2018/01/01", 1514764800),
		buildDemo("yyyy.MM.dd", "2018.01.01", 1514764800),
		buildDemo("yyyy年MM月dd日", "2018年01月01日", 1514764800),
		//zh-04
		buildDemo("MM-dd hh:mm", "03-01 12:10", 1519906200),
		buildDemo("MM/dd hh:mm", "03-01 12:10", 1519906200),
		buildDemo("MM年dd月 hh:mm", "03-01 12:10", 1519906200),
		//us-01
		buildDemo("May. 23. 1998", "May. 23. 1998", 895881600),
		buildDemo("May. 23th. 1998", "May. 23th. 1998", 895881600),
		buildDemo("July 23. 1998", "July 23. 1998", 901152000),
		buildDemo("July 23th. 1998", "July 23th. 1998", 901152000),
		buildDemo("May. 23, 1998", "May. 23, 1998", 895881600),
		buildDemo("May. 23th, 1998", "May. 23th, 1998", 895881600),
		buildDemo("July 23, 1998", "July 23, 1998", 901152000),
		buildDemo("July 23th, 1998", "July 23th, 1998", 901152000),
		buildDemo("23 July, 1998", "July 23, 1998", 901152000),
		buildDemo("23th July, 1998", "July 23th, 1998", 901152000),
		buildDemo("23 July. 1998", "July 23, 1998", 901152000),
		buildDemo("23th Jul. 1998", "July 23th, 1998", 901152000),
		buildDemo("1st Jul. 1998", "July 1st, 1998", 899251200),
		//us-02
		buildDemo("dd-MM-yyyy hh:mm:ss", "01-01-2018 12:12:12", 1514808732),
		buildDemo("dd.MM.yyyy hh:mm:ss", "01-01-2018 12:12:12", 1514808732),
		buildDemo("dd MM yyyy hh:mm:ss", "01-01-2018 12:12:12", 1514808732),
		//us-03
		buildDemo("dd-MM-yyyy hh:mm", "2018-01-01 12:12", 1514808720),
		buildDemo("dd/MM/yyyy hh:mm", "2018/01/01 12:12", 1514808720),
		buildDemo("dd.MM.yyyy hh:mm", "2018.01.01 12:12", 1514808720),
		//us-04
		buildDemo("dd-MM-yyyy'", "01-01-2018", 1514764800),
		buildDemo("dd/MM/yyyy", "01/01/2018", 1514764800),
		buildDemo("dd.MM.yyyy", "01.01.2018", 1514764800),
		//co-01
		// buildDemo("seconds ago", "1s ago", time.Now().Add(time.Duration(1)*time.Second).Unix()),
		// buildDemo("mins ago", "1min ago", time.Now().Add(time.Duration(1)*time.Minute).Unix()),
		// buildDemo("houres ago", "1 hours ago", time.Now().Add(time.Duration(1)*time.Hour).Unix()),
		// buildDemo("day ago", "1 day ago", time.Now().Add(time.Duration(24)*time.Hour).Unix()),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindTime(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
