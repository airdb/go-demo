package main

import (
  "fmt"
  "io/ioutil"
)


func main() {
  file, error := ioutil.TempFile("/tmp","golang-file")
  defer file.Close()
  if error != nil {
    fmt.Println("create file failed")
    return
  }

  for {
  file.WriteString("Hello world")
  }
  //filedata, _ := ioutil.ReadFile(file.Name())
  //fmt.Println(string(filedata))

}

