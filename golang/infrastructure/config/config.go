package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port int
}

func (config *Config) loadEnv() *Config {
	err := godotenv.Load()
	if os.IsNotExist(err) {
		return nil
	}
	return config
}

//loadSpecification allows to decode all environment variables
//into an config defined specification struct
func (config *Config) LoadConfig() *Config {

	spec := &Config{}
	err := envconfig.Process("MY_THERESA", spec)
	if err != nil {
		panic(err)
	}
	config = spec
	return config
}

//Init initializes the configuration struct
func Init() *Config {
	configInstance := &Config{}
	config := configInstance.
		loadEnv().
		LoadConfig()

	return config
}
