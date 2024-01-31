package main

import (
	"fmt"
	"os"
)

func main() {
	for _, msg := range os.Args[1:] {
		fmt.Printf("%v", msg)
	}
}
