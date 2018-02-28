package main

import "fmt"

func main() {
  messages := make(chan string, 4)

  messages <- "channel buffer 1"
  messages <- "channel buffer 2"
  messages <- "channel buffer 3"
  messages <- "channel buffer 4"

  fmt.Println(<-messages)
  fmt.Println(<-messages)
  fmt.Println(<-messages)
  //fmt.Println(<-messages)
}
