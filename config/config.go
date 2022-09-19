package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App      AppMeta
	DBConfig MySQLMeta
}

type AppMeta struct {
	Listener string `yaml:"listener"`
	Mode     string `yaml:"mode"`
}

type MySQLMeta struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Minpool  int    `yaml:"minpool"`
	Maxpool  int    `yaml:"maxpool"`
}

func LoadConfig(path string) (config *Config) {
	b, err := os.ReadFile(path)
    if err != nil {
        panic("Failed to load config file")
    }

    err = yaml.Unmarshal(b, &config)
    if err != nil {
        panic("Failed to load config file")
    }

    return
}
