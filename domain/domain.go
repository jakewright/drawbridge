package domain

import (
	"net/url"
)

// API represents a single upstream API to proxy requests to
type API struct {
	Name             string `yaml:"name"`
	Prefix           string `yaml:"prefix"`
	UpstreamURL      *URL   `yaml:"upstream_url"`
	AllowCrossOrigin bool   `yaml:"allow_cross_origin"`
}

// URL wraps the net/url.URL type to provide an UnmarshalYAML function
type URL struct {
	*url.URL
}

// UnmarshalYAML tells the parser how to unmarshal a string to type Url
func (u *URL) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string

	// Unmarshal the field into a string
	if err := unmarshal(&s); err != nil {
		return err
	}

	// Parse the URL string into a url.URL type
	target, err := url.Parse(s)
	if err != nil {
		return err
	}

	// Set the anonymous field on u
	u.URL = target

	// Return no errors
	return nil
}
