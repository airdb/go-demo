package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

const (
	numElements = 10000000
)

func timeGC() {
	t := time.Now()
	runtime.GC()
	fmt.Printf("gc took: %s\n", time.Since(t))
}

var pointers = map[string]int{}

func main() {
	for i := 0; i < 10000000; i++ {
		pointers[strconv.Itoa(i)] = i
	}

	for {
		timeGC()
		time.Sleep(1 * time.Second)
	}
}
