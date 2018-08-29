package config

import (
	"log"
	"sync"

	"github.com/tokopedia/tdk/go/config"
)

type Config struct {
	Appname string
}

var cfg Config
var once = sync.Once{}

func GetConfig() Config {
	once.Do(func() {
		err := config.Read(&cfg, "config/punyasaya.{TKPENV}.yaml", "/etc/punyasaya/punyasaya.{TKPENV}.yaml")
		if err != nil {
			log.Fatal(err)
		}
	})
	return cfg
}
