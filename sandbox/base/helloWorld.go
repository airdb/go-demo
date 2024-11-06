package main

import "fmt"


func Add(a int , b int ) int {
  return a +b
}
func main() {
  a := 111
  b := 111
  c  := Add( a, b)
  fmt.Println( c ) 
  // fmt.Println("Hello MT.Qomolangma!")
}
