package config

import (
	"fmt"
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
	//log.Printf("Path is %s", path)
	configDir := path.Dir(file)
	configFile := path.Base(file)
	extension := strings.Split(path.Ext(file), ".")
	configExt := extension[len(extension)-1]

	fmt.Printf("Configdir is %s\n", configDir)
	fmt.Printf("ConfigFile is %s\n", configFile)

	viper.AddConfigPath(configDir)

	// Set the file name and type of the configurations file
	viper.SetConfigFile(configFile)
	viper.SetConfigType(configExt)

	// Enable Viper to read Environment Variables
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

	return config, nil
}
