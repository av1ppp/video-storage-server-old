package config

type DatabaseConfig struct {
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	DatabaseName string `yaml:"dbname"`
	SSLMode      bool   `yaml:"sslmode"`
}
