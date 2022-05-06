package main

import (
	"fmt"
	"os"

	"github.com/arduino/arduino-cli/arduino/serialutils"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("reset-arduino <port>\n")
		os.Exit(1)
	}
	port := os.Args[1]
	err := serialutils.TouchSerialPortAt1200bps(port)
	if err != nil {
		fmt.Printf("cannot reset: %s", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
