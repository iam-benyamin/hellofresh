package config

import (
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func Load(configPath string) Config {
	k := koanf.New(".")

	err := k.Load(file.Provider(configPath), yaml.Parser())
	if err != nil {
		panic("load config file fail")
	}

	err = k.Load(env.Provider("HELLOFRESH_", ".", func(s string) string {
		str := strings.Replace(strings.ToLower(strings.TrimPrefix(s, "HELLOFRESH_")), "_", ".", -1)

		return strings.Replace(str, "..", "_", -1)
	}), nil)
	if err != nil {
		panic("load config env fail")
	}

	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		panic(err)
	}
	return cfg
}
