package app

import (
	"fmt"
	"io/ioutil"

	"github.com/tapokshot/shelld1core/log"
	database "github.com/tapokshot/shelld1core/store"
	"gopkg.in/yaml.v2"
)

//todo support json, env file config

// Config for httpServer application
type Config struct {
	// BindAddr bind port for httpServer
	BindAddr string `yaml:"bind_addr"`
	// LogLevel log level (debug, info .etc)
	LoggerCfg *log.Config `yaml:"logger"`
	// PostgresDB config for postgresql
	PostgresDB *database.PostgresDBConfig `yaml:"postgres-db"`
}

// CreateConfig parse file from configPath and create Config
func CreateConfig(configPath string) (*Config, error) {
	config := &Config{}
	err := parseYaml(configPath, config)
	if err != nil {
		return nil, err
	}
	return config, err
}

func parseYaml(configPath string, config *Config) error {
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) String() string {
	return fmt.Sprintf("%#+v", c)
}
