package main

import (
    "io"
    "log"
    "net/http"
    "time"
    "strconv"
)


var port string = "8002"

var mux map[string]func(http.ResponseWriter, *http.Request)
func main() {
    server := http.Server{
        Addr:           ":"+port,
        Handler:        &MyHandle{},
        ReadTimeout:    6 * time.Second,
        WriteTimeout:   6 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    mux = make(map[string]func(http.ResponseWriter, *http.Request))
    mux["/"] = api_help
    mux["/api"] = http_api
    mux["/hello"] = hello
    mux["/bye"] = bye
    mux["/mdb"] = bye
    mux["/service"] = bye
    mux["/monitor"] = bye


    err := server.ListenAndServe()
    if err != nil {
        log.Fatal(err)
    }
}

type MyHandle struct{}

func (*MyHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if h, ok := mux[r.URL.String()]; ok {
        h(w, r)
    }

    str := ""
    if "/status" == r.URL.String() {
       str += status()
    } else if "/api/status" == r.URL.String() {
       str += status()
    }
    str += "\n"
    io.WriteString(w, str)
}



func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "hello module")
}

func bye(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "bye module")
}

func http_api(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "/api  http_api")
}

func status() string {
  //return "status function return"
  var status string = "OK"
  return status
  return strconv.Itoa(0)
}

//func status() (string){
//  return "status"
//}

func api_help(w http.ResponseWriter, req *http.Request) {
  // fmt.Println(req.Host, req.URL.Path, req.URL.Path[1:])
  msg := ""
  msg += "\nhttp-api is http-server which exposed the client status and using information.\n"
  msg += "Writeby Dean, email is osfun@qq.com.\n\n"
  msg += "Version: 1.0.0\n\n"
  msg += "API Usage:\n  curl http://localhost"+port+"/api\n"
  msg += "\n===============================\n"
  msg += "status\n"
  msg += "service\n"
  msg += "mdb\n"
  msg += "==============END=============="

  w.Write([]byte(msg))
}

