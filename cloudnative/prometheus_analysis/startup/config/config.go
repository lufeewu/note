package config

import (
	"io/ioutil"

	"github.com/caarlos0/env"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	WebConfig WebConfig `yaml:"web"`
	original  string
}

type WebConfig struct {
	Value string `yaml:"value" env:"Value" envDefault:"bbbbb"`
}

func Load(s string) (*Config, error) {
	cfg := &Config{}

	// default from env
	env.Parse(cfg)

	err := yaml.UnmarshalStrict([]byte(s), cfg)
	if err != nil {
		return nil, err
	}

	cfg.original = s
	return cfg, nil
}

func LoadFile(filename string) (*Config, error) {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	cfg, err := Load(string(content))
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
