package conf

import (
	"cchw2/pkg/ninja"
	"cchw2/pkg/redis"
	"fmt"
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
)

// Config
// struct type of app configs.
type Config struct {
	Redis redis.Config `koanf:"redis"`
	Ninja ninja.Ninja  `koanf:"ninja"`
}

// Load
// loading app configs.
func Load() Config {
	var instance Config

	k := koanf.New(".")

	// load default
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		_ = fmt.Errorf("error loading deafult: %v\n", err)
	}

	// load configs file
	if err := k.Load(file.Provider("config.yaml"), yaml.Parser()); err != nil {
		_ = fmt.Errorf("error loading config.yaml file: %v\n", err)
		fmt.Println(err)
	}

	// unmarshalling
	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshalling config: %v\n", err)
	}

	return instance
}

func Default() Config {
	return Config{
		Redis: redis.Config{
			Host:    "",
			Timeout: 0,
		},
		Ninja: ninja.Ninja{
			APIkey: "",
		},
	}
}
