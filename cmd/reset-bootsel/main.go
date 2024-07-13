package main

import (
	"fmt"
	"os"

	serialutils "github.com/arduino/go-serial-utils"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("%s <port>\n", os.Args[0])
		os.Exit(1)
	}
	port := os.Args[1]
	err := serialutils.Touch1200bps(port)
	if err != nil {
		fmt.Printf("cannot reset: %s", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
