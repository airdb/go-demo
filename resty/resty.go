package main

import (
	"log"

	"github.com/go-resty/resty/v2"
)

type Resp struct {
	Origin string `json:"origin"`
	Url    string `json:"url"`
}

func main() {

	client := resty.New()
	var resp Resp
	_, err := client.R().SetResult(&resp).
		Get("https://httpbin.org/get")
	if err != nil {
		panic(err)
	}

	log.Println(resp)
}
