package main

import (
	"fmt"

	"github.com/zelcion/focusmode/configuration"
	"github.com/zelcion/focusmode/hosts"
)

func main() {
	// This below has no utility, but ensures that we have valid configuration
	_, err := configuration.InitializeConfig()

	if err != nil {
		fmt.Printf("Error initializing configuration: %v\n", err)
		panic(err)
	}

	// cErr := hosts.CreateExampleHostsFile()
	// if cErr != nil {
	// 	fmt.Printf("Error creating example hosts file: %v\n", cErr)
	// 	panic(cErr)
	// }

	_, err = hosts.CanEditHostsFile()
	if err != nil {
		fmt.Printf("FocusMode must run with elevated permissions to edit the hosts file. \n")
		fmt.Printf("Error checking hosts file permissions: %v\n", err)
		panic(err)
	}

	// // Load hosts file to test addFmAnnotationToFile function
	// fmt.Println("Loading hosts file to debug addFmAnnotationToFile...")
	// hostFile, err := hosts.LoadHostsFile()
	// if err != nil {
	// 	fmt.Printf("Error loading hosts file: %v\n", err)
	// } else {
	// 	fmt.Printf("Successfully loaded hosts file with %d entries\n", len(hostFile.Entries))
	// }

	// hosts.SetBlockForDomain("www.facebook.com", true)
	hosts.RestoreBackup()
	fmt.Println("Focusmode initialized successfully.")
}
