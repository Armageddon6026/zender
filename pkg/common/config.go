package common

import (
	"encoding/json"
	"io"
	"os"
)

type SystemConfig struct {
	Port    string   `json:"port,omitempty"`
	WebRoot string   `json:"webroot,omitempty"`
	Db      Dbconfig `json:"db,omitempty"`
}

type Dbconfig struct {
	Addr     string `json:"addr,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
}

func GetSystemConfig() (*SystemConfig, error) {
	config := new(SystemConfig)
	err := readJsonFromFile(SYSTEM_CONFIG_PATH, config)
	if err != nil {
		return nil, err
	}

	setDbConfigByEnv(config)

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

func setDbConfigByEnv(c *SystemConfig) {
	if DB_ADDR != "" {
		c.Db.Addr = DB_ADDR
	}
	if DB_PASSWORD != "" {
		c.Db.Password = DB_PASSWORD
	}
	if DB_USER != "" {
		c.Db.User = DB_USER
	}
}
