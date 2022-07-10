package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ConfigDb struct {
	Type     string
	Name     string
	Host     string
	Username string
	Password string
	Port     string
	Database string
	Driver   string
}

type Config struct {
	Port string
	Dbs  []ConfigDb
}

func ReadConfig(path string) (*Config, error) {

	var config *Config

	if path == "" {
		path = "../config.json"
	}

	// Open our jsonFile
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully Opened", path)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &config)

	return config, err
}
