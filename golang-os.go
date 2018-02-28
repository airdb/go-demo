package main

import (
  "fmt"
  "os"
)


func main() {

  fmt.Println(os.Chdir("/tmp"))
  dir, _ := os.Getwd()
  fmt.Println("current working dir is: ", dir)


  PS1 := os.Getenv("PS1")
  fmt.Println("env PS1 is: ", PS1)

  fmt.Println(  os.Getuid(), os.Getgid(), os.Geteuid(), os.Getegid() )

  g, _ := os.Getgroups()
  fmt.Println(g)

  fmt.Println("swap size is: ",os.Getpagesize())

  fmt.Println("process id is: ", os.Getppid())

  filemode, _ := os.Stat("/bin/ls")
  fmt.Println("/bin/ls :",filemode.Mode())

  // err := os.Chmod("widuu.go", 0777)
  // if err!=nil{
  //   fmt.Println("修改文件权限失败")
  // }

  data, _ := os.Hostname()
  fmt.Println(data)


  environ := os.Environ()
  fmt.Println(environ[39])

}
