package main

import (
  "fmt"
  "os/exec"
)

func main() {
  res, err := exec.LookPath("php")
  if err != nil {
    fmt.Println("not install php")
  }
  fmt.Println(res)
  fmt.Println()
  fmt.Println("==========================")


  argv := []string{"-l"}
  // c := exec.Command("ls", argv...)
  c := exec.Command("echo", argv...)
  domain := "a.map.gtimg.com"
  a := "http://lb.wsd.com/lb/interface?skey=&operator=deanlzhang&interface_name=get_domain_info&interface_params=%7B%20%22f_domain%22:%22" + domain +"%22%20%7D"
  c = exec.Command("curl", a)
  d, _ := c.Output()
  fmt.Println("exec ls command: ")
  fmt.Println(string(d))
  fmt.Println()
  fmt.Println("==========================")


  result, error := exec.Command("date").Output()
  if error != nil {
    fmt.Println(result)
  }
  fmt.Println(string(result))

}
