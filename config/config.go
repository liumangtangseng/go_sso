package config

type Serve struct {
	Jwt Jwt `yaml:"jwt"`
	Mysql Mysql `yaml:"mysql"`
	System System `yaml:"system"`
	Loggers Loggers `yaml:"loggers"`
}
