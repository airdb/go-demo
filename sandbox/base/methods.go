package main

import "fmt"

type rect struct {
  width, height int
}


func (r *rect) area() int {
  return r.width * r.height
}

func (r *rect) perim() int {
  return 2 * (r.width + r.height)
}

func main() {
  r := rect{width: 6, height: 8}
  fmt.Println("area: ", r.area() )
  fmt.Println("perim: ", r.perim() )

  rp := &r
  fmt.Println("point area: ", rp.area() )
  fmt.Println("point perim: ", rp.perim() )

}
