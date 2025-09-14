package main

import (
	"fmt"

	"github.com/zelcion/focusmode/configuration"
)

func main() {
	_, err := configuration.InitializeConfig()

	if err != nil {
		fmt.Printf("Error initializing configuration: %v\n", err)
		panic(err)
	}

	fmt.Printf("Configuration loaded successfully \n")
}
