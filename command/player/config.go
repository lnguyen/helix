package player

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//Config for player
type Config struct {
	DataDir string `json:"data_dir"`
	Name    string `json:"name"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

//DefaultConfig if none is passed in
func DefaultConfig() *Config {
	return &Config{
		DataDir: ".",
		Host:    "0.0.0.0",
		Name:    "Player-1",
		Port:    8821,
	}
}

//ReadConfig opens up file for reading
func ReadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Error reading '%s': %s", path, err)
	}
	return DecodeConfig(f)
}

//DecodeConfig from reader
func DecodeConfig(r io.Reader) (*Config, error) {
	var result Config
	rawConfig, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(rawConfig, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

//MergeConfig merge config files together
func MergeConfig(a, b *Config) *Config {
	var result Config = *a

	if b.DataDir != "" {
		result.DataDir = b.DataDir
	}
	if b.Host != "" {
		result.Host = b.Host
	}
	if b.Name != "" {
		result.Name = b.Name
	}
	if b.Port != 0 {
		result.Port = b.Port
	}

	return &result
}
