package settings

import (
	_ "embed"
	"gopkg.in/yaml.v3"
)

//go:embed settings.yaml
var settingsFile []byte

type DataBaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Settings struct {
	Port string         `yaml:"port"`
	DB   DataBaseConfig `yaml:"database"`
}

func New() (*Settings, error) {

	var s Settings

	if err := yaml.Unmarshal(settingsFile, &s); err != nil {
		return nil, err
	}

	return &s, nil
}
