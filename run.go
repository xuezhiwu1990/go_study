package main

import (
        "fmt"
        "strings"
        "strconv"
		"flag"
		"net/url"
		"math/rand"
		"time"
		"os"
)

type resource struct {
  url string
  target string
  start int
  end int
}

func ruleResource () []resource {
  var res [] resource
  r1 := resource{
    url : "http://www.go.com/" ,
    target: "",
    start: 0,
    end:0,
  }

  r2:= resource{
    url : "http://www.go.com/list/{$id}.html" ,
    target: "{$id}",
    start: 1,
    end:21,
  }

  r3 := resource{
    url : "http://www.go.com/movie/{$id}.html" ,
    target: "{$id}",
    start: 1,
    end:15308,
  }
  res = append(append(append(res,r1),r2),r3)
  return res
}

func buildUrl( res []resource) []string {
  var list []string
  for _,resItem := range res {
    if len(resItem.target) == 0 {
      list = append(list ,resItem.url)
    }else{
      for i:=resItem.start;i<=resItem.end;i++{
          urlStr :=  strings.Replace( resItem.url , resItem.target , strconv.Itoa(i) , -1 )
          list = append(list ,urlStr)
      }
    }
  }
  return list
}

func makelog( current , refer , ua  string) string {
 u := url.Values{}
 u.Set("time","1")
 u.Set("url",current)
 u.Set("refer",refer)
 u.Set("ua",ua)
 paramsStr := u.Encode()

 logTemplate := "192.168.47.1 - - [03/Jun/2018:12:21:01 +0800] 'OPTIONS /dig?{$paramsStr}' '{$ua}' '-'";
 //logTemplate := "{$paramsStr}{$ua}";
 log := strings.Replace(logTemplate , "{$paramsStr}",paramsStr,-1)
 log = strings.Replace(log , "{$ua}",ua,-1)
 return log
}

func randInt(min ,max int) int {
	fmt.Println(min,max)
	r := rand.New( rand.NewSource(time.Now().UnixNano() ))
	if min > max {
		return max
	}
	return r.Intn(max-min) + min
}

func main () {
  total := flag.Int( "total" , 100 ,"how many")
  filePath := flag.String("filePath","D:\\Go_WorkSpace\\dig_go.log" ,"D:\\Go_WorkSpace\\dig_go.log")
  flag.Parse()

  //需要构造出真实的url集合
  res := ruleResource()
  list := buildUrl( res )
  fmt.Println(list)

  //按照要求,生成$total 行日志内容
	logStr := ""
  for i:= 0;i<=*total;i++ {
	  currentUrl := list[ randInt(0, len(list)-i) ]
	  referUrl := list[ randInt(0, len(list)-i) ]
	  ua := "Mozilla/5.0 (Windows NT 10.0.10586; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2950.5 Safari/537.36"
      logStr = logStr + makelog(currentUrl,referUrl,ua ) + "\n"
  }
  fd,_ := os.OpenFile( *filePath , os.O_RDWR | os.O_APPEND , 0644)
  fd.Write([]byte(logStr))
  fd.Close()
  fmt.Println("done.\n")
}
