package config 

import (
	"os" // gives access to finding os stuff like finding the home directory
	"path/filepath" // filepath.Join idomatic method
	"encoding/json" // Decode
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}

	path := filepath.Join(homeDir, configFileName)
	
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	
	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return Config{}, err
	}
	return config, nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	if err := write(*cfg); err != nil {
		return err
	}
	return nil
}

func write(cfg Config) error {
	homeDir, err := os.UserHomeDir()
        if err != nil { 
                return err
        }

        path := filepath.Join(homeDir, configFileName)
	
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(cfg); err != nil {
		return err
	}

	return nil
}
