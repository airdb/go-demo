package main

import (
	"bytes"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"sync/atomic"
)

func main() {
	var reqID uint64
	http.HandleFunc("/plaintext", func(w http.ResponseWriter, req *http.Request) {
		id := atomic.AddUint64(&reqID, 1)
		buf := new(bytes.Buffer)
		for i := 0; i < 100; i++ {
			buf.WriteString(strconv.Itoa(int(id)))
			buf.WriteByte(' ')
			buf.WriteString(req.RemoteAddr)
		}
		w.Write(buf.Bytes())
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
