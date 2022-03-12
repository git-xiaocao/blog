package config

import (
	_ "embed"
	"gopkg.in/yaml.v3"
)

type config struct {
	Database database
}

type database struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int16  `yaml:"port"`
	DBName   string `yaml:"db_name"`
}

//go:embed config.yaml
var configYamlFile []byte

var configYaml config

func init() {
	_ = yaml.Unmarshal(configYamlFile, &configYaml)
}

func Config() *config {
	return &configYaml
}
