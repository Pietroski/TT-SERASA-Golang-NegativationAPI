package util

import "github.com/spf13/viper"

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
		return
	}

	err = viper.Unmarshal(&config)
	return
}
