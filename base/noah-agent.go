package main

import (
  "fmt"
  "time"
  "io/ioutil"
  "log"
  "net/http"
  "math/rand"
)

func main() {

  // dead loop
  count := 0
  for {
    count += 1


    //log.Println(time.Now().Unix())
    DOMAIN := "http://install.ksyun.com/"

    timestamp := time.Now().Unix()
    url := "webfoot?"
    url += fmt.Sprintf("taskid=%d&timestamp=%d",count,timestamp)

    interface_url := DOMAIN + url
    user_agent := "osfunBot"
    referer := "noah"

    client := &http.Client{}
    req, err := http.NewRequest("GET", interface_url, nil)
    if err != nil {
      log.Fatalln(err)
    }

    req.Header.Set("Content-Type", "text/plain")
    req.Header.Set("User-Agent", user_agent)
    req.Header.Set("Referer", referer )

    req.Header.Set("X-Custom-Header", "myvalue")

    resp, err := client.Do(req)
    if err != nil {
      log.Fatalln(err)
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Fatalln(err)
      log.Println(string(body))
    }

    log.Println("count:", count)

    // Sleep time
    rand.Seed(time.Now().UnixNano())
    time.Sleep(time.Second * time.Duration(rand.Intn(60)) )
    //time.Sleep(time.Second  * rand.Intn(10))
  }

}
