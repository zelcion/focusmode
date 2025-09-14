package hosts

import (
	"fmt"
	"os"

	"github.com/zelcion/focusmode/constants"
)

func SetBlockForDomain(domain string, active bool) error {
	hostFile, err := LoadHostsFile()
	if err != nil {
		return err
	}

	found := false
	for i, entry := range hostFile.Entries {
		if entry.Domain == domain {
			entry.Active = active
			hostFile.Entries[i] = entry
			found = true
			break
		}
	}

	if !found {
		newEntry := HostEntry{
			RedirectTo:  constants.REDIRECT_IP,
			Domain:      domain,
			Active:      active,
			IsFmManaged: true,
		}
		hostFile.Entries[domain+constants.REDIRECT_IP] = newEntry
	}

	return hostFile.Save()
}

func setBlockAllDomains(status bool) error {
	hostFile, err := LoadHostsFile()
	if err != nil {
		return err
	}

	changed := false
	for i, entry := range hostFile.Entries {
		if entry.IsFmManaged && entry.Active == !status {
			entry.Active = status
			hostFile.Entries[i] = entry
			changed = true
		}
	}

	if changed {
		return hostFile.Save()
	}

	return nil
}

func BlockAllDomains() error {
	return setBlockAllDomains(true)
}

func UnblockAllDomains() error {
	return setBlockAllDomains(false)
}

func RestoreBackup() error {
	hostsFilePath, err := getHostsFilePath()
	if err != nil {
		return err
	}

	backupPath := hostsFilePath + constants.HOSTS_FILE_BACKUP_EXTENSION

	input, err := os.ReadFile(backupPath)
	if err != nil {
		return fmt.Errorf("failed to read backup file: %v", err)
	}

	err = os.WriteFile(hostsFilePath, input, 0644)
	if err != nil {
		return fmt.Errorf("failed to restore hosts file from backup: %v", err)
	}

	return nil
}
