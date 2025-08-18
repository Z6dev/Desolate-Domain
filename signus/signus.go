package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Channel to receive OS signals
	sigChan := make(chan os.Signal, 1)
	i := 0

	// Notify on SIGINT (Ctrl+C)
	signal.Notify(sigChan, syscall.SIGINT)

	fmt.Println("Program running... Press Ctrl+C to make it meow!")

	for {
		<-sigChan
		fmt.Println("\b\bMeow! 🐱")
		i++

		if i >= 3 {
			os.Exit(0)
		}
	}
}
