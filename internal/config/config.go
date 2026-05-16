package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

var filePath string

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	filePath = filepath.Join(home, ".gatorconfig.json")
}

func Read() Config {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		log.Fatal(err)
	}
	return config
}

func (c *Config) SetUser(username string) {
	c.CurrentUserName = username
	c.Write()
}

func (c Config) Write() {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(c); err != nil {
		log.Fatal(err)
	}
}
