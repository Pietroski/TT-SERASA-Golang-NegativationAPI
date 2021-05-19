package util

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	Config iConfig = &SConfig{}
)

type iConfig interface {
	LoadConfig(path string) (config SConfig, err error)
}

// SConfig stores all the application's configurations variables.
// The values are read by viper from a config file or environment variables.
type SConfig struct {
	DBDriver                   string `mapstructure:"DB_DRIVER"`
	DBDataSourceName           string `mapstructure:"DB_DATA_SOURCE_NAME"`
	NegativationsServerAddress string `mapstructure:"NEGATIVATIONS_SERVER_ADDRESS"`
	LegacyServerAddress        string `mapstructure:"LEGACY_SERVER_ADDRESS"`
}

// LoadConfig reads configuration from a file or environment variable.
func (c *SConfig) LoadConfig(path string) (config SConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)

		// This block of code was introduced due to docker
		// Viper was unable to Automatically read from env somehow.
		// TODO: remove this piece of code
		viper.SetDefault("DB_DRIVER", "postgres")
		viper.SetDefault("DB_DATA_SOURCE_NAME", "postgresql://serasa:serasa_psql@localhost:5432/tt_serasa?sslmode=disable")
		viper.SetDefault("NEGATIVATIONS_SERVER_ADDRESS", "localhost:8008")
		viper.SetDefault("LEGACY_SERVER_ADDRESS", "localhost:8009")

		//return
	}

	err = viper.Unmarshal(&config)
	return
}
