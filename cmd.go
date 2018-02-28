package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input") //输入

	var out bytes.Buffer
	cmd.Stdout = &out //输出

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("in all caps: %q\n", out.String())

}
