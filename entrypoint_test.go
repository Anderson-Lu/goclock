package entrypoint

import (
	core "go_time_finder/core"
	"reflect"
	"testing"
)

type a struct {
	source string
}

type demo struct{		
		name string
		args a
		want []core.FinderResult		
}

func buildDemo(format,args string,result int64) demo{
	r:=demo{
		name: "[TESTING] "+format,
		args: a{source:args},
		want: []core.FinderResult{
			core.FinderResult{
				SourceStr: args,
				ResultStr: args,
				ResultUTC: result,
			},
		},
	}
	return r
}


func TestFindTime(t *testing.T) {
	
	tests := []demo{
		//zh-01
		buildDemo("yyyy-MM-dd hh:mm:ss","2018-01-01 12:12:12",1514808732),
		buildDemo("yyyy/MM/dd hh:mm:ss","2018/01/01 12:12:12",1514808732),
		buildDemo("yyyy.MM.dd hh:mm:ss","2018.01.01 12:12:12",1514808732),
		buildDemo("yyyy年MM月dd日 hh:mm:ss","2018年01月01日 12:12:12",1514808732),
		buildDemo("yyyy年MM月dd日 hh时mm分ss秒","2018年01月01日 12时12分12秒",1514808732),	
		buildDemo("yyyy年MM月dd日 hh時mm分ss秒","2018年01月01日 12時12分12秒",1514808732),
		//zh-02
		buildDemo("yyyy-MM-dd hh:mm","2018-01-01 12:12",1514808720),
		buildDemo("yyyy/MM/dd hh:mm","2018/01/01 12:12",1514808720),
		buildDemo("yyyy.MM.dd hh:mm","2018.01.01 12:12",1514808720),
		buildDemo("yyyy年MM月dd日 hh:mm","2018年01月01日 12:12",1514808720),
		buildDemo("yyyy年MM月dd日 hh时mm分","2018年01月01日 12时12分",1514808720),	
		buildDemo("yyyy年MM月dd日 hh時mm分","2018年01月01日 12時12分",1514808720),
		//zh-03
		buildDemo("yyyy-MM-dd","2018-01-01",1514764800),
		buildDemo("yyyy/MM/dd","2018/01/01",1514764800),
		buildDemo("yyyy.MM.dd","2018.01.01",1514764800),
		buildDemo("yyyy年MM月dd日","2018年01月01日",1514764800),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindTime(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
