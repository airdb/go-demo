package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func main() {
    var data struct {
        //var code int
        Data []struct {
            country    string
            country_id string
            area  string
            area_id  int
            region  string
region_id  int
            city  string
city_id  int
            county  string
county_id  int
            isp  string
isp_id  int
        }
    }

   url := "http://ip.taobao.com/service/getIpInfo.php?ip=182.118.53.124"

    r, _ := http.Get(url)
    defer r.Body.Close()

    dec := json.NewDecoder(r.Body)
    dec.Decode(&data)

    for _, item := range data.Data{
        //fmt.Printf("%s = %d\n", item.Name, item.Count)
       fmt.Println("aaa",item.country_id)
    }

}
