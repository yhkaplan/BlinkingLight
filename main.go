package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
	"time"
)

var (
	pin = rpio.Pin(4)
)

func main() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio when finished
	defer rpio.Close()

	// Set pin to "output mode"
	pin.Output()

	// Toggle
	// NOTE: ending on an even number results in the light
	// remaining on
	for i := 0; i < 101; i++ {
		pin.Toggle()
		time.Sleep(time.Second)
	}
}
