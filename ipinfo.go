package main

import (
   "os"
    "fmt"
    "time"
    "encoding/json"
    curl "github.com/andelf/go-curl"
)

func main() {
    easy := curl.EasyInit()
    defer easy.Cleanup()

    //easy.Setopt(curl.OPT_URL, "http://install.ksyun.com/")

    count := 0
    for {
      count += 1
	  ip := os.Args[1:][0]
      url := "http://ip.taobao.com/service/getIpInfo.php?ip=" + ip
		  // fmt.Println(url)
      easy.Setopt(curl.OPT_URL, url)

      // make a callback function
      fooTest := func (buf []byte, userdata interface{}) bool {
          // println("DEBUG: content=>", string(buf))

          //json.Unmarshal([]byte(string(buf)), &res)
// 		  j2 := make(map[string]interface{})
// 		  json.Unmarshal([]byte(string(buf)), &j2)
// 		  j3 := make(map[string]interface{})
// 		  a := j2["data"](map[string]interface{})
// 		  fmt.Println(j2["data"])
// 		  fmt.Println(a)

          var f  interface{}
		  json.Unmarshal([]byte(string(buf)), &f)
		  d := f.(map[string]interface{})
		  c := d["data"].(map[string]interface{})

		  fmt.Println(c["ip"],c["country"])
          return true
      }

      easy.Setopt(curl.OPT_WRITEFUNCTION, fooTest)
      if err := easy.Perform(); err != nil {
          fmt.Printf("ERROR: %v\n", err)
      }

      break
      time.Sleep(time.Second)
    }

}
