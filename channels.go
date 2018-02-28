package main

import "fmt"

func main() {

  // create a channel
  messages := make(chan string)

  go func()  { 

    messages <- "ping" 
  } ()

  msg := <-messages
  fmt.Println(msg)

}
