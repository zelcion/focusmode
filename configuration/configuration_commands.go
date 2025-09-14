package configuration

func (c *Configuration) SetDomainBlock(domain string, active bool) error {
	config, err := InitializeConfig()
	if err != nil {
		return err
	}
	configPath := GetDefaultConfigPath()

	config.setDomainBlock(domain, active)
	config.Save(configPath)
	return nil
}
