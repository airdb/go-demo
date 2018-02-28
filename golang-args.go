package main

import "os"
import "fmt"

func main() {
  fmt.Println("I am", os.Args[1:][0])
}
