package main

import (
  "fmt"
  //"io/ioutil"
  "os"
  "time"
)


func main() {
  filename := "/tmp/golang-write-file"
  // f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
  f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    panic(err)
  }

  defer f.Close()

  for {
    text := "Hello World\n"
    if _, err = f.WriteString(text); err != nil {
      fmt.Println("aaaa")
      panic(err)
    }
    time.Sleep(time.Second * 1)
  }
}
