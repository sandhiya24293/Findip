package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"Rentmatics_App/Logger"
)

var (
	log = Logger.NewLogger("rent")
)

func main() {

	getConfig()

	if ok := Serve(); ok {
		sig := make(chan os.Signal)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	forever:
		for {
			select {
			case s := <-sig:
				fmt.Println("Signal ", s, " received, stopping")
				break forever
			}
		}
		fmt.Println("Server stopped")

	}

}
