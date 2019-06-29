package config

import (
	"fmt"

	"github.com/cc2k19/go-tin/storage"

	"github.com/cc2k19/go-tin/server"
	"github.com/spf13/viper"
)

// InputValidator should be implemented by types that need input validation check. Mainly settings types
type InputValidator interface {
	Validate() error
}

// ConfigFile describes the name, path and the format of the file to be used to load the configuration in the env
type ConfigFile struct {
	Name     string `description:"name of the configuration file"`
	Location string `description:"location of the configuration file"`
	Format   string `description:"extension of the configuration file"`
}

// Settings is used to setup go-tin
type Settings struct {
	Storage *storage.Settings
	Server  *server.Settings
}

// Validate validates the server settings
func (s *Settings) Validate() error {
	validatable := []InputValidator{s.Server, s.Storage}
	for _, v := range validatable {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// DefaultSettings returns the default values for configuring go-tin
func DefaultSettings() *Settings {
	return &Settings{
		Server:  server.DefaultSettings(),
		Storage: storage.DefaultSettings(),
	}
}

// DefaultConfigFile holds the default config file properties
func DefaultConfigFile() ConfigFile {
	return ConfigFile{
		Name:     "application",
		Location: ".",
		Format:   "yml",
	}
}

// New creates a configuration from the environment
func New() (*Settings, error) {
	config := DefaultSettings()
	configFile := DefaultConfigFile()

	v := viper.New()

	v.AddConfigPath(configFile.Location)
	v.SetConfigName(configFile.Name)
	v.SetConfigType(configFile.Format)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("could not read configuration cfg: %s", err)
		}
	}

	if err := v.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("error loading configuration: %s", err)
	}

	return config, nil
}
