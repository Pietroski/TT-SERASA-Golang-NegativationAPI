package util

import "github.com/spf13/viper"

var (
	Config iConfig = &sConfig{}
)

type iConfig interface {
	LoadConfig(path string) (config sConfig, err error)
}

// sConfig stores all the application's configurations variables.
// The values are read by viper from a config file or environment variables.
type sConfig struct {
	DBDriver                    string `mapstructure:"DB_DRIVER"`
	DBDataSourceName            string `mapstructure:"DB_DATA_SOURCE_NAME"`
	NeggativationsServerAddress string `mapstructure:"NEGATIVATIONS_SERVER_ADDRESS"`
}

// LoadConfig reads configuration from a file or environment variable.
func (c *sConfig) LoadConfig(path string) (config sConfig, err error) {
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
