package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

func ParseFile(filename string, conf *Config) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(conf)
	if err != nil {
		return err
	}

	return nil
}

func ParseEnv(conf *Config) error {
	return envconfig.Process("", conf)
}
