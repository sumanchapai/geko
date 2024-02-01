package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "0.0.0"

func main() {
	// Version
	versionPtr := flag.Bool("version", false, "version")
	flag.Parse()
	if *versionPtr {
		fmt.Println(version)
		return
	}
	// Normal
	for i, msg := range os.Args[1:] {
		if i == 0 {
			fmt.Printf("%v", msg)
		} else {
			fmt.Printf(" %v", msg)
		}
	}
	fmt.Println()
}
