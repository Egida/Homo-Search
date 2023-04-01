package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ShodanApiKey   string
	SearchSettings struct {
		Pages int
		Query struct {
			Text   []string
			Os     string
			Hash   int
			HasSSL bool
			Port   int
			Net    string
			Org    string
			Tag    string
			Region int
		}
	}
}

func ReadConfig() *Config {

	config, err := os.ReadFile("./config.json")

	if err != nil {
		fmt.Println("homo search: can't find the config file")
		os.Exit(0)
	}

	var _config Config
	json.Unmarshal(config, &_config)

	return &_config
}
