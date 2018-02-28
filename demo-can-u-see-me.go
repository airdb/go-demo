package main

import "fmt"
// import "strings"

func main() {
  var arr = [6]int{1, 2, 3}   

  var msg string
  i := 6
  if i > 5 {
    i = 5
  } else {
    msg = "Can u see me?"
  }

  fmt.Println( arr[i], msg, i )
}
