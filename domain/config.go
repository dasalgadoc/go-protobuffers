package domain

import "errors"

type Config struct {
	Port     string
	Host     string
	Network  string
	Database string
}

func (c Config) ConfigErrors() error {
	if c.Port == "" {
		return errors.New("Config: port is required")
	}
	if c.Host == "" {
		errors.New("Config: host is required")
	}
	if c.Network == "" {
		return errors.New("Config: network is required")
	}
	if c.Database == "" {
		return errors.New("Config: database is required")
	}
	return nil
}
