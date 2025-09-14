package hosts

import (
	"fmt"
	"os"
)

const EXAMPLE_HOSTS_FILE_CONTENT = `# FocusMode edited this file 
# Original file backup is created with .backup extension.
127.0.0.1       localhost
::1             localhost
`

func CreateExampleHostsFile() error {
	hostsFilePath, err := getHostsFilePath()
	if err != nil {
		return err
	}

	fmt.Printf("Creating example hosts file at %s\n", hostsFilePath)
	return os.WriteFile(hostsFilePath, []byte(EXAMPLE_HOSTS_FILE_CONTENT), 0644)
}
