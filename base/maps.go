package main
import "fmt"
func main() {
  m := make(map[string]int)
  m["k1"] = 7
  m["k2"] = 71

  fmt.Println("map:", m)
  fmt.Println("k1 value:", m["k1"])

}
