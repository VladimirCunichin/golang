package config

import (
	"strings"

	"bitbucket.org/VladimirCunichin/golang/src/master/calendar/internal/logger"
	"github.com/spf13/viper"
)

type Config struct {
	Log        Log        `json:"log" mapstructure:"log"`
	HttpListen HttpListen `json:"http" mapstructure:"http"`
	DBConfig   DBConfig   `json:"db_config" mapstructure:"db"`
}

type Log struct {
	LogFile  string `json:"log_file" mapstructure:"log_file"`
	LogLevel string `json:"log_level" mapstructure:"log_level"`
}

type HttpListen struct {
	Ip   string `json:"ip" mapstructure:"ip"`
	Port string `json:"port" mapstructure:"port"`
}

type DBConfig struct {
	User     string `json:"user" mapstructure:"user"`
	Password string `json:"password" mapstructure:"password"`
	Host     string `json:"host" mapstructure:"host"`
	Port     string `json:"port" mapstructure:"port"`
	Database string `json:"database" mapstructure:"database"`
}

func GetConfigFromFile(filePath string) *Config {
	viper.SetConfigFile(filePath)
	if err := viper.ReadInConfig(); err != nil {
		logger.Fatal("Couldn't read configuration file", "error", err)
	}
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	var C Config
	err := viper.Unmarshal(&C)
	if err != nil {
		logger.Fatal("error during unmarshall", "error", err)
	}
	return &C
}
