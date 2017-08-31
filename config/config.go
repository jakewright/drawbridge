package config

import (
    "drawbridge/domain"
    "gopkg.in/yaml.v2"
    "io/ioutil"
)

type Configuration struct {
    Apis map[string]domain.Api
}

func GetApiConfig() (c Configuration) {
    // Read the configuration file
    b, err := ioutil.ReadFile("/config/config.yaml")
    if err != nil { panic(err) }

    // Unmarshal the YAML into the Configuration struct
    err = yaml.Unmarshal(b, &c)
    if err != nil {
        panic(err)
    }

    // Return the Configuration struct
    return
}
