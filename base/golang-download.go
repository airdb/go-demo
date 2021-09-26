package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
)

func downloadFromUrl(url string) {
    tokens := strings.Split(url, "/")
    fileName := tokens[len(tokens)-1]
    fmt.Println("Downloading", url, "to", fileName)

    // TODO: check file existence first with io.IsExist
    output, err := os.Create(fileName)
    if err != nil {
        fmt.Println("Error while creating", fileName, "-", err)
        return
    }
    defer output.Close()

    response, err := http.Get(url)
    if err != nil {
        fmt.Println("Error while downloading", url, "-", err)
        return
    }
    defer response.Body.Close()

    n, err := io.Copy(output, response.Body)
    if err != nil {
        fmt.Println("Error while downloading", url, "-", err)
        return
    }

    fmt.Println(n, "bytes downloaded.")
}

func main() {
        url := "http://install.ksyun.com/sys_tools.tgz"
        downloadFromUrl(url)
}
