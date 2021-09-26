package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

func main() {
	fmt.Printf("hello, world\n")
	fmt.Println("hello, world ", time.Now())
	fmt.Println(math.Pi)
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1:])
	}
}
