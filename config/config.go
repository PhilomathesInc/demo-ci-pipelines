package config

import (
	"io/ioutil"
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

const configFile = "config.yaml"

type Config struct {
	Server server `yaml:"server"`
	Logger *zap.Logger
}

type server struct {
	Protocol string `yaml:"protocol"`
	Hostname string `yaml:"hostname"`
	Port     string `yaml:"port"`
	Timeout  string `yaml:"timeout"`
}

func New() (*Config, error) {
	c := &Config{}

	// Read configs from config file
	f, err := os.Open(configFile)
	if err != nil {
		return c, err
	}
	defer f.Close()
	if b, err := ioutil.ReadAll(f); err != nil {
		return c, err
	} else if len(b) != 0 {
		yaml.Unmarshal(b, c)
	}

	// Initializing logger
	if c.Logger, err = zap.NewDevelopment(); err != nil {
		return c, err
	}

	return c, nil
}
