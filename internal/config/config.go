package config

type Config struct {
	Env     string `yaml:"env"`
	MongoDB struct {
		Host       string `yaml:"host"`
		Port       uint   `yaml:"port"`
		Username   string `yaml:"username"`
		Password   string `yaml:"password"`
		Database   string `yaml:"database"`
		Collection string `yaml:"collection"`
	} `yaml:"mongodb"`
}
