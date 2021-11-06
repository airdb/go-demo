package main

import (
	"log"
	"strings"
	"time"
)

func main() {
	siteLoc := " +0800"
	date := "2021/09/27 12:30:03" + siteLoc

	layout := "2006/01/02 15:04:05 -0700"

	t, err := time.Parse(layout, strings.TrimSpace(date))
	if err != nil {
		log.Panic(err)
	}

	log.Println("time", t.Unix())
	log.Println("time", t)
}
