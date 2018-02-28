package main

  import (
  "fmt"
  )

  type LesssAdder interface {
      Less(b Integer) bool
      Add(b Integer)
  }

  type Integer int

  func (a Integer) Less(b Integer) bool {
      return a < b
  }

  func (a *Integer) Add(b Integer) {
      *a += b
  }

  func main() {

      var a Integer = 1
      var b LesssAdder = &a
      fmt.Println(b)

      //var c LesssAdder = a
      //Error:Integer does not implement LesssAdder 
      //(Add method has pointer receiver)
  }
