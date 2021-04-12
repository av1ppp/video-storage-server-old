package config

type ServerConfig struct {
	Host string `yaml:"host", envconfig:"SERVER_HOST"`
	Port int    `yaml:"port", envconfig:"SERVER_PORT"`
}
