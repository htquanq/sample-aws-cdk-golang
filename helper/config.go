package helper

import (
	. "example-app-go/interfaces"

	"gopkg.in/yaml.v3"
)

func Config(yamlFile []byte) AppConfig {
	var config AppConfig
	yaml.Unmarshal(yamlFile, &config)

	return config
}
