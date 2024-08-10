package common

import (
	"encoding/json"
	"io"
	"os"
)

type SystemConfig struct {
	Addr    string   `json:"addr,omitempty"`
	WebRoot string   `json:"webroot,omitempty"`
	Db      Dbconfig `json:"db,omitempty"`
}

type Dbconfig struct {
	Addr     string `json:"addr,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	DBName   string `json:"dbname,omitempty"`
}

func GetSystemConfig() (*SystemConfig, error) {
	config := new(SystemConfig)
	err := readJsonFromFile(SYSTEM_CONFIG_PATH, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func readJsonFromFile(path string, c *SystemConfig) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	value, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(value, c)
	if err != nil {
		return err
	}
	return nil
}
