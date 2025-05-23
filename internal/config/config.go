package config

import "time"

type Config struct {
	Env     string `yaml:"env"`
	MongoDB struct {
		Host        string        `yaml:"host"`
		Port        int           `yaml:"port"`
		Username    string        `yaml:"username"`
		Password    string        `yaml:"password"`
		Database    string        `yaml:"database"`
		Collection  string        `yaml:"collection"`
		ConnTimeout time.Duration `yaml:"conntimeout"`
	} `yaml:"mongodb"`
}
