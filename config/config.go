package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port         string `yaml:"port"`
		ReadTimeout  int    `yaml:"readTimeout"`
		WriteTimeout int    `yaml:"write_timeout"`
		IdleTimeout  int    `yaml:"idle_timeout"`
	} `yaml:"server"`
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
