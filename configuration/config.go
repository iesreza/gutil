package configuration

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Configuration struct {
	app  string
	path []string
}

func getInstance() *Configuration {
	cfg := Configuration{}
	cfg.app = os.Args[0]
	cfg.path = []string{"/etc/" + cfg.app + "/", "$HOME/.", "."}

	return &cfg
}

func (cfg Configuration) load() {
	viper.SetConfigName(cfg.app + ".conf")
	for _, path := range cfg.path {
		viper.AddConfigPath(path)
	}
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
