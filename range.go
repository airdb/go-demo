package main

import "fmt"

func mapStringString() {
 a := map[string]string{ "f_lb_inner_ip": "f_lb_region", "f_rs_ip": "f_rs_region", }
  for k,v := range a {
    fmt.Printf("k: %s, v: %s\n", k, v )
  }
}

func mapStringInt() (string, int) {
  a := map[string]int { "f_lb_inner_ip": 2, "f_rs_ip": 11 }
  sum := 0
  stringSum := ""
  for k,v := range a {
    fmt.Printf("k: %s, v: %d\n", k, v )
    sum += v
    stringSum += k
  }
  return stringSum,sum
}

func mapIntString() {
  a := map[int]string { 2:"f_lb_inner_ip", 4:"f_rs_ip" }
  for k,v := range a {
    fmt.Printf("k: %d, v: %s\n", k, v )
  }
}

func main() {
  // a := []string{"aaa", "bbb", "ccc"}
  // for k,v := range a {
  //   fmt.Println( aa )
  //   fmt.Printf("k: %d, v: %s\n", k, v )
  // }
  
  keyReturn,testReturn := mapStringInt()
  fmt.Println( testReturn,keyReturn )

}
