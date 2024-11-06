package main

import (
  "fmt"
  "math"
)


type geometry interface {
  area()  float64
  perim()  float64
}


type rect struct{
  width, height float64
}


type circle struct {
  radius float64
}


func (r rect) area() float64 {
  return r.width * r.height
}

func (r rect) perim() float64 {
  return 2 * (r.width + r.height)
}


func measure( g geometry) {
  fmt.Println(g)
  fmt.Println(g.area())
  fmt.Println(g.perim())
}


func main() {
  r := rect{width: 6, height: 8 }
  measure(r)
  fmt.Println(math.Pi)
}
