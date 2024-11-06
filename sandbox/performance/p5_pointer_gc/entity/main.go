package main

import (
	"fmt"
	"runtime"
	"time"
)

func timeGC() {
	t := time.Now()
	runtime.GC()
	fmt.Printf("gc took: %s\n", time.Since(t))
}

type Entity struct {
	A int
	B float64
}

var entities = map[Entity]int{}

func main() {
	for i := 0; i < 10000000; i++ {
		entities[Entity{
			A: i,
			B: float64(i),
		}] = i
	}

	for {
		timeGC()
		time.Sleep(1 * time.Second)
	}
}
