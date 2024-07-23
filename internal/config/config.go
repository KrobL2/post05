package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	HTTPServer `yaml:"http_server"`
	DB         `yaml:"db"`
}

type HTTPServer struct {
	Port         string        `yaml:"port"`
	Timeout      time.Duration `yaml:"timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

func LoadConfig(filename string) (*Config, error) {
	config := &Config{}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(data))

	err = yaml.UnmarshalStrict(data, config)
	if err != nil {
		return nil, err
	}

	fmt.Println(config)

	return config, nil
}
