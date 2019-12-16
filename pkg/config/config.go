package config

import (
	"time"

	"github.com/BurntSushi/toml"
)

type Config struct {
	AutoRia AutoRia `toml:"autoria"`
}

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

type AutoRia struct {
	Period Duration `toml:"period"`
	ApiKey string   `toml:"api_key"`
}

// New reads application configuration from specified file path.
func New(path string) (*Config, error) {
	var config Config

	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
