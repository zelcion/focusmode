package commands

import (
	"fmt"

	"github.com/zelcion/focusmode/configuration"
	"github.com/zelcion/focusmode/hosts"
)

func BlockDomain(domain string) error {
	config, err := configuration.InitializeConfig()
	if err != nil {
		return fmt.Errorf("failed to initialize configuration: %v", err)
	}

	// Update the configuration
	err = config.SetDomainBlock(domain, true)
	if err != nil {
		return fmt.Errorf("failed to update configuration: %v", err)
	}

	// Update the hosts file
	return hosts.SetBlockForDomain(domain, true)
}
