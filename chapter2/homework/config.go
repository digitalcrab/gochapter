package homework

import (
	"github.com/jessevdk/go-flags"
)

type Config struct {
	HTTPListen string `long:"listen" description:"Listen address" default:":8888" required:"true"`
	GoogleKey  string `long:"key" description:"Google Maps API Key" required:"true"`
}

// NewConfig return the instance of config structure or error if so
func NewConfig() (*Config, error) {
	var c Config

	if _, err := flags.NewParser(&c, flags.HelpFlag|flags.PassDoubleDash).Parse(); err != nil {
		return nil, err
	}

	return &c, nil
}
