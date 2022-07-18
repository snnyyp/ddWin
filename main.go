package main

import (
	"os"
	"fmt"
)


func main() {
	stdout, err := os.OpenFile("conerr$", os.O_RDWR, 0000)
	defer stdout.Close()
	fmt.Println(err)
	stdout.WriteString("Hello, world! from CONOUT$")
}
