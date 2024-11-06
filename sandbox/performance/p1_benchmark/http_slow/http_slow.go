package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync/atomic"
)

func main() {
	var reqID uint64
	http.HandleFunc("/plaintext", func(w http.ResponseWriter, req *http.Request) {
		id := atomic.AddUint64(&reqID, 1)
		var s string
		for i := 0; i < 100; i++ {
			s += fmt.Sprint(id, " ", req.RemoteAddr)
		}
		w.Write([]byte(s))
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
