package config

import (
	"io/ioutil"

	"github.com/jakewright/drawbridge/domain"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	APIs map[string]domain.Api
}

func Load(filename string) (*Configuration, error) {
	// Read the configuration file
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Unmarshal the YAML into the Configuration struct
	var c Configuration
	err = yaml.Unmarshal(b, &c)
	if err != nil {
		return nil, err
	}

	// Return the Configuration struct
	return &c, nil
}
