package main

import (
    "crypto/md5"
    "fmt"
    "io"
    "os"
)

func main() {
    testFile := "./glog.go"
    file, inerr := os.Open(testFile)
    if inerr == nil {
        md5h := md5.New()
        io.Copy(md5h, file)
        fmt.Printf("%x  %s\n", md5h.Sum([]byte("")), testFile) //md5
    }
}
