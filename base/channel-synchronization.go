package main

import "fmt"
import "time"

func worker(done chan bool) {
  fmt.Print("working...")
  time.Sleep(time.Second)
  fmt.Println("done")

  done <- true
}


func main() {
  done := make(chan bool, 2)
  for i :=0; i< 10; i++{
  go worker(done)
  }
  for i :=0; i< 100; i++{
  <-done
  }
  
}
