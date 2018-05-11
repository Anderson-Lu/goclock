GO TIME FINDER
---

![author](https://img.shields.io/badge/author-anderson--lu-yellow.svg?longCache=true&style=flat-square)
![build](https://img.shields.io/badge/build-passing-green.svg?longCache=true&style=flat-square)

GO-TIME-FINDER is a text time identification library written by Golang.

Install
---

```shell
$ go get github.com/Anderson-Lu/go_time_finder
```

Usage
---

```golang
import core "github.com/Anderson-Lu/go_time_finder/core"

func main() {
  t := core.FindTime('今天是2018年1月1日')
  fmt.Println(t)
}
```


Support Formats
---

|format|example|result|
|:-|:-|:-|
|`yyyy-MM-dd hh:mm:ss`|`2018-01-01 12:12:12`|1514808732|
|`yyyy/MM/dd hh:mm:ss`|`2018/01/01 12:12:12`|1514808732|
|`yyyy.MM.dd hh:mm:ss`|`2018.01.01 12:12:12`|1514808732|
|`yyyy年MM月dd日 hh:mm:ss`|`2018年01月01日 12:12:12`|1514808732|
|`yyyy年MM月dd日 hh时mm分ss秒`|`2018年01月01日 12时12分12秒`|1514808732|
|`yyyy年MM月dd日 hh時mm分ss秒`|`2018年01月01日 12時12分12秒`|1514808732|
|`yyyy-MM-dd hh:mm`|`2018-01-01 12:12`|1514808720|
|`yyyy/MM/dd hh:mm`|`2018/01/01 12:12`|1514808720|
|`yyyy.MM.dd hh:mm`|`2018.01.01 12:12`|1514808720|
|`yyyy年MM月dd日 hh:mm`|`2018年01月01日 12:12`|1514808720|
|`yyyy年MM月dd日 hh时mm分`|`2018年01月01日 12时12分`|1514808720|
|`yyyy年MM月dd日 hh時mm分`|`2018年01月01日 12時12分`|1514808720|
|`yyyy-MM-dd`|`2018-01-01`|1514764800|
|`yyyy/MM/dd`|`2018/01/01`|1514764800|
|`yyyy.MM.dd`|`2018.01.01`|1514764800|
|`yyyy年MM月dd日`|`2018年01月01日`|1514764800|
|`MM-dd hh:mm`|`03-01 12:10`|1519906200|
|`MM/dd hh:mm`|`03-01 12:10`|1519906200|
|`MM年dd月 hh:mm`|`03-01 12:10`|1519906200|
|`May. 23. 1998`|`May. 23. 1998`|895881600|
|`May. 23th. 1998`|`May. 23th. 1998`|895881600|
|`July 23. 1998`|`July 23. 1998`|901152000|
|`July 23th. 1998`|`July 23th. 1998`|901152000|
|`May. 23 1998`|`May. 23 1998`|895881600|
|`May. 23th, 1998`|`May. 23th, 1998`|895881600|
|`July 23, 1998`|`July 23, 1998`|901152000|
|`July 23th, 1998`|`July 23th, 1998`|901152000|
|`23 July, 1998`|`July 23, 1998`|901152000|
|`23th July 998`|`23th July 1998`|901152000|
|`23 July. 1998`|`23 July. 1998`|901152000|
|`23th Jul. 1998`|`23th Jul. 1998`|901152000|
|`1st Jul. 1998`|`July 1st. 1998`|899251200|
|`dd-MM-yyyy hh:mm:ss`|`01-01-2018 12:12:12`|1514808732|
|`dd.MM.yyyy hh:mm:ss`|`01-01-2018 12:12:12`|1514808732|
|`dd MM yyyy hh:mm:ss`|`01-01-2018 12:12:12`|1514808732|
|`dd-MM-yyyy hh:mm`|`2018-01-01 12:12`|1514808720|
|`dd/MM/yyyy hh:mm`|`2018/01/01 12:12`|1514808720|
|`dd.MM.yyyy hh:mm`|`2018.01.01 12:12`|1514808720|
|`dd-MM-yyyy'`|`01-01-2018`|1514764800|
|`dd/MM/yyyy`|`01/01/2018`|1514764800|
|`dd.MM.yyyy`|`01.01.2018`|1514764800|
|`seconds ago`|`1s ago`|-|
|`mins ago`|`1min ago`|-|
|`houres ago`|`1 hours ago`|-|
|`day ago`|`1 day ago`|-|
