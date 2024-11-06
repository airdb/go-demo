package main

import "fmt"

type person struct {
  name string
  age int
}


func main() {

  s := person{ name: "Sean", age: 50}
  fmt.Println("type:", s.name, s.age)
}
