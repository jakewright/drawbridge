package config

import (
	"io/ioutil"

	"github.com/jakewright/drawbridge/domain"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Apis map[string]domain.Api
}

func LoadConfig() (c Configuration) {
	// Read the configuration file
	b, err := ioutil.ReadFile("/config/config.yaml")
	if err != nil {
		panic(err)
	}

	// Unmarshal the YAML into the Configuration struct
	err = yaml.Unmarshal(b, &c)
	if err != nil {
		panic(err)
	}

	// Return the Configuration struct
	return
}
