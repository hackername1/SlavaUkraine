package main

import (
	"fmt"
	"time"
)

func main() {
	// Print "Слава Україні!" and "Героям слава!" in turn with a pause per second
	for {
		fmt.Println("Слава Україні!")
		time.Sleep(time.Second)
		fmt.Println("Героям слава!")
		time.Sleep(time.Second)
	}
}
