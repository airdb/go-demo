package main

import (
  "fmt"
  "os"
)


func main() {
  var fd string = "/tmp/defer.txt"
  fmt.Println("file is: ", fd)
  f := createFile(fd)
  defer closeFile(f)
  writeFile(f)
}


func createFile(p string) *os.File {
  fmt.Println("create")
  f, err := os.Create(p)

  if err != nil {
    panic(err)
  }

  return f
}


func writeFile(f *os.File) {
  fmt.Println("writing")
  fmt.Fprintln(f,"data")
}


func closeFile(f *os.File) {
  fmt.Println("closing")
  f.Close()
}
