package server

import "fmt"

// Settings type to be loaded from the environment
type Settings struct {
	Port int `mapstructure:"port" description:"port of the server"`
}

// DefaultSettings returns the default values for configuring the server
func DefaultSettings() *Settings {
	return &Settings{
		Port: 8080,
	}
}

// Validate validates the server settings
func (s *Settings) Validate() error {
	if s.Port == 0 {
		return fmt.Errorf("validate Settings: Port missing")
	}
	return nil
}
