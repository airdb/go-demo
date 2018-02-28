package main

import "fmt"

func foo() {
  defer fmt.Println("World")
  fmt.Print("Hello ")
}


func sum(x,y int, c chan int) {
  c <- x + y
}


func main() {
  foo()
  c := make (chan int)
  go sum(24,18,c)
  fmt.Println(<-c)
}



// 通过 Channel 的概念, fmt.Println(<-c) 会阻塞代码执行,
// 而 go sum(24,18,c) 启动的 goroutine 则继续执行,
// 等到 c <- x + y 往 Channel 里发送了数据,
// fmt.Println(<-c) 接收到了数据, 才继续执行代码,
// 这样就比较方便能用同步的逻辑写异步代码, Node 也方便, 但是不够优雅
