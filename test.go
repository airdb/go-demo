package main

import (
  "fmt"
  "os"
)

func main() {
cmdToRun := "/home/zhanglei16/dean/golang/Writefile"
args := []string{""}
procAttr := new(os.ProcAttr)
procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
fmt.Println("Parent Process id: ", os.Getppid())
if process, err := os.StartProcess(cmdToRun, args, procAttr); err != nil {
    fmt.Printf("ERROR Unable to run %s: %s", cmdToRun, err.Error())
} else {
    fmt.Printf("%s running as pid %d", cmdToRun, process.Pid)
}
}
