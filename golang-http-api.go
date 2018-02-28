package main

import (
  "net/http"
  "fmt"
  "encoding/json"
)


// scheme://[userinfo@]host/path[?query][#fragment]
func api_help(w http.ResponseWriter, req *http.Request) {
  fmt.Println(req.Host, req.URL.Path, req.URL.Path[1:])
  w.Write([]byte("\nhttp-api is http-server which exposed the client status and using information.\n"))
  w.Write([]byte("Writeby Dean, email is osfun@qq.com.\n\n"))
  w.Write([]byte("Version: 1.0.0\n\n"))
  w.Write([]byte("API Usage:\n  curl http://localhost"+port+"/api\n"))
  w.Write([]byte("\n===============================\n"))
  w.Write([]byte("status\n"))
  w.Write([]byte("service\n"))
  w.Write([]byte("mdb\n"))
  w.Write([]byte("==============END==============\n"))
}

func api_args(w http.ResponseWriter, req *http.Request) {
  w.Write([]byte("status    port    service\n"))
}

func api_status(w http.ResponseWriter, req *http.Request) {
  w.Write([]byte("status: "+status+"\n"))
}

func api_service(w http.ResponseWriter, req *http.Request) {
  w.Write([]byte("nginx.noah.all\nweb.noah.all\n"))
}

func api_mdb(w http.ResponseWriter, req *http.Request) {
  w.Write([]byte("SN: 8MYQ342\n"))
}

func api_json(w http.ResponseWriter, req *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=utf-8")
  myItems := []string{"item1", "item2", "item3"}
  a, _ := json.Marshal(myItems)
  w.Write(a)
}

var port string = "8002"
var status string = "OK"

func main() {

  // route
  http.HandleFunc("/", api_help)
  http.HandleFunc("/api", api_args)
  http.HandleFunc("/api/status", api_status)
  http.HandleFunc("/api/service", api_service)
  http.HandleFunc("/api/mdb", api_mdb)
  http.HandleFunc("/api/json", api_json)
  //http.HandleFunc("/login", )
  //http.ListenAndServe(":8002",nil)
  // http.ListenAndServe(port, nil)

  // listen
  err := http.ListenAndServe(":"+port, nil)
  if err != nil {
    fmt.Println("ListenAndServe: ", err)
  }

}
