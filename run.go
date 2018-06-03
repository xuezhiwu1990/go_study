package main

import (
       "fmt"
       "strings"
       "strconv"
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

//func makelog( current , refer , ua  string) string {
//  u : = url.Values{}
//  u.Set("time","1")
//  u.Set("url",current)
//  u.Set("refer",refer)
//  u.Set("ua",ua)
//  //paramsStr := utf16.Encode()
//}

func main () {
  //total := flag.Int( "total" , 100 ,"how many")
  //filePath := flag.String("filePath","./dig_go.log")
  //flag.Parse()

  //需要构造出真实的url集合
  res := ruleResource()
  list := buildUrl( res )
  fmt.Println(list)

  //按照要求,生成$total 行日志内容
  //for i:= 0;i<=*total;i++ {
  //  logStr := makelog(currentUrl , referUrl ,ua )
  //}

}
