package configuration

import (
	"os"

	"github.com/zelcion/focusmode/constants"

	"gopkg.in/yaml.v3"
)

type DomainEntry struct {
	Domain string `yaml:"domain"`
	Active bool   `yaml:"active"`
}

type Configuration struct {
	Version            string        `yaml:"version"`
	BlockedDomains     []DomainEntry `yaml:"blocked_domains"`
	BlockedExecutables []string      `yaml:"blocked_executables"`
	Active             bool          `yaml:"active"`
}

func LoadConfiguration(path string) (*Configuration, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var config Configuration

	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Configuration) Save(path string) error {
	data, err := yaml.Marshal(c)

	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func NewEmptyConfiguration() *Configuration {
	return &Configuration{
		Version:            constants.CURRENT_VERSION,
		BlockedDomains:     []DomainEntry{},
		BlockedExecutables: []string{},
		Active:             false,
	}
}

func (c *Configuration) setDomainBlock(domain string, active bool) {
	exists := false
	for i, entry := range c.BlockedDomains {
		if entry.Domain == domain {
			c.BlockedDomains[i].Active = active
			exists = true
			break
		}
	}
	if !exists {
		c.BlockedDomains = append(c.BlockedDomains, DomainEntry{Domain: domain, Active: active})
	}
}
