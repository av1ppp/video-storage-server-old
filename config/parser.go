package config

import (
	"os"

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
