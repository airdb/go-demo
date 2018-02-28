package main

import "fmt"
import "os"

var _ = os.Exit

var  aaa int = 1 + 1

func Add(a int , b int ) int {
  fmt.Println( aaa )
  return a +b
}

func main() {
  a := 111
  b := 111
  c  := Add( a, b)
  // os.Exit( 1 )
  fmt.Println( c ) 
  // fmt.Println("Hello MT.Qomolangma!")
}
