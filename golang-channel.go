package main

import (
"fmt"
"time"
)

func worker(start chan bool, index int){
  <- start  // get a value from channel
  fmt.Println("This is Worker: ", index)
}


func main() {
  start := make(chan bool)
  for i := 1; i <= 100; i++ {
    go worker( start, i)
  }
 close(start)

 time.Sleep(time.Second * 5)
 select {}
}
