package plugin

import (
	"fmt"
	"sync"

	"github.com/jakewright/muxinator"
	"github.com/mitchellh/mapstructure"
)

var (
	l = sync.RWMutex{}
	m = map[string]Plugin{}
)

type Plugin interface {
	Middleware(config map[string]interface{}) (muxinator.Middleware, error)
}

func Find(name string) (Plugin, error) {
	l.RLock()
	defer l.RUnlock()

	plugin, ok := m[name]
	if !ok {
		return nil, fmt.Errorf("plugin '%s' not found", name)
	}

	return plugin, nil
}

func RegisterPlugin(name string, plugin Plugin) {
	l.Lock()
	defer l.Unlock()

	m[name] = plugin
}

type Validator interface {
	Validate() error
}

func DecodeConfig(cfg map[string]interface{}, opts Validator) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result: opts,
	})
	if err != nil {
		return err
	}

	if err := decoder.Decode(cfg); err != nil {
		return err
	}

	return opts.Validate()
}
