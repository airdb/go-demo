package main

import (
    "encoding/json"
    "fmt"
)

type Server struct {
    ServerName string
    ServerIP   string
}

type Serverslice struct {
    Servers []Server
}

func main() {
    var s Serverslice
    str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
    json.Unmarshal([]byte(str), &s)
    fmt.Println(s)

    // str1 := `{"code":0,"data":{"country":"\u4e2d\u56fd","country_id":"CN","area":"\u534e\u4e2d","area_id":"400000","region":"\u6cb3\u5357\u7701","region_id":"410000","city":"\u90d1\u5dde\u5e02","city_id":"410100","county":"","county_id":"-1","isp":"\u8054\u901a","isp_id":"100026","ip":"182.118.53.124"}}`
}
