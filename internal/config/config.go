package config

import (
	"log"
	"path"
	"strings"

	"github.com/spf13/viper"
)

// Configuration for the Cortex Client
// So far, the only mandatory value is
// ApiKey, since BaseUrl has a default
// value.
type Config struct {
	BaseUrl string `mapstructure:"BASE_URL"`
	ApiKey  string `mapstructure:"API_KEY"`
}

// GetConfig returns a Config object to retreive information
// for both Request creation and Client instantiation
func GetConfig(file string) (*Config, error) {
	// Set the path to look for the configurations file
	configDir := path.Dir(file)
	configFile := path.Base(file)
	extension := strings.Split(path.Ext(file), ".")
	configExt := extension[len(extension)-1]

	viper.AddConfigPath(configDir)

	// Set the file name and type of the configurations file
	viper.SetConfigFile(configFile)
	viper.SetConfigType(configExt)

	// Enable Viper to read Environment Variables
	//   TODO: Not working as expected, still requires
	//   a `config.env` file since it's igonring the
	//   environment variables.
	//   there are some discussions around it in the upstream:
	//   https://github.com/spf13/viper/issues/761
	viper.AutomaticEnv()

	// Read config
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
		return nil, err
	}

	// Unmarshall values into target Config object
	config := &Config{}
	err := viper.Unmarshal(config)
	if err != nil {
		log.Printf("Error when unmarshall config, %s", err)
		return nil, err
	}

	log.Printf("Configuration has been successfully loaded")
	return config, nil
}
