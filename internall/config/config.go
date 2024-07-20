package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port         string        `yaml:"port"`
		ReadTimeout  time.Duration `yaml:"read_timeout"`
		WriteTimeout time.Duration `yaml:"write_timeout"`
		IdleTimeout  time.Duration `yaml:"idle_timeout"`
	} `yaml:"http_server"`
}

func LoadConfig(filename string) (*Config, error) {
	config := &Config{}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// fmt.Print(string(data))

	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	// fmt.Print(config)

	return config, nil
}
