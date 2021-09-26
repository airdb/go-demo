package main

import (
    "fmt"
    "time"
    curl "github.com/andelf/go-curl"
)

func main() {
    easy := curl.EasyInit()
    defer easy.Cleanup()

    //easy.Setopt(curl.OPT_URL, "http://install.ksyun.com/")

    count := 0
    for {
      count += 1
      easy.Setopt(curl.OPT_URL, "http://install.ksyun.com")

      // make a callback function
      fooTest := func (buf []byte, userdata interface{}) bool {
          println(count, "DEBUG: size=>", len(buf))
          // println("DEBUG: content=>", string(buf))
          return true
      }
      easy.Setopt(curl.OPT_WRITEFUNCTION, fooTest)
      if err := easy.Perform(); err != nil {
          fmt.Printf("ERROR: %v\n", err)
      }

      time.Sleep(time.Second)
    }

}
