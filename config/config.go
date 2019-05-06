package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Configuration represents the config loaded from the YAML file
type Configuration struct {
	Port int             `yaml:"port"`
	APIs map[string]*API `yaml:"apis"`
}

// API is a single API definition
type API struct {
	Name             string   `yaml:"name"`
	Prefix           string   `yaml:"prefix"`
	UpstreamURL      string   `yaml:"upstream_url"`
	AllowCrossOrigin bool     `yaml:"allow_cross_origin"`
	Plugins          []Plugin `yaml:"plugins"`
}

// Plugin is the configuration for a particular plugin
type Plugin struct {
	Name    string                 `yaml:"name"`
	Enabled bool                   `yaml:"enabled"`
	Config  map[string]interface{} `yaml:"config"`
}

// Load reads a YAML file and returns a Configuration struct
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
