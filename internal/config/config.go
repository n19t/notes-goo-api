package config

import (
	"io/ioutil"
	"log"
	"notes-goo-api/internal/env"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type ConfigDb struct {
	Type     string `yaml:"type"`
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Driver   string `yaml:"driver"`
}

type Config struct {
	Port string     `yaml:"port"`
	Dbs  []ConfigDb `yaml:"dbs"`
}

var (
	path = env.GetEnvVariable("CONFIG_FILE")
)

func ReadConfig() (*Config, error) {
	var config Config

	if path == "" {
		path = "./config.yaml"
	}

	filename, _ := filepath.Abs(path)
	byteValue, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatal(err)
	}

	return &config, nil
}
