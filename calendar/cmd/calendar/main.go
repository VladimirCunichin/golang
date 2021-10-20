package main

import (
	"flag"

	"github.com/spf13/viper"
	"github.com/vladimircunichin/golang/calendar/internal/logger"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "config/config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()
	vi := viper.New()
	vi.SetConfigFile(configFile)
	vi.ReadInConfig()
	logger.Init(vi.GetString("log_level"), vi.GetString("log_file"))

}
