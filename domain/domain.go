package domain

import (
    "net/url"
)

type Api struct {
    Name string
    Prefix string
    UpstreamUrl *Url `yaml:"upstream_url"`
}

type Url struct {
    *url.URL
}

// Tell the parser how to unmarshal a string to type Url
func (u *Url) UnmarshalYAML(unmarshal func(interface{}) error) error {
    var s string

    // Unmarshal the field into a string
    if err := unmarshal(&s); err != nil {
        return err
    }

    // Parse the URL string into a url.URL type
    target, err := url.Parse(s)
    if err != nil { return err }

    // Set the anonymous field on u
    u.URL = target

    // Return no errors
    return nil
}
