package environment

import (
	"github.com/gofiber/fiber/v2/log"

	"github.com/spf13/viper"
)

var (
	EnvKey      = "ENV"
	ServicePort = "SERVICE_PORT"
	DsnKey      = "DB_DSN"
)

func New(dirDepth uint) {
	viper.SetDefault("ENV", "dev")
	env := viper.GetString(EnvKey)
	viper.SetConfigName("app-" + env)
	viper.SetConfigType("yaml")
	log.Info("ENV: " + env)

	var configDir string
	if dirDepth == 0 {
		configDir = "."
	} else {
		configDir = ".."
		for i := uint(1); i < dirDepth; i++ {
			configDir += "/.."
		}
	}

	viper.AddConfigPath(configDir)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

func GetString(key string) string {
	if !viper.IsSet(key) {
		panic("failed to get environment key: " + key)
	}

	return viper.GetString(key)
}

func GetStrings(key string) []string {
	return viper.GetStringSlice(key)
}

func GetInt(key string) int {
	if !viper.IsSet(key) {
		panic("failed to get environment key: " + key)
	}

	return viper.GetInt(key)
}

func GetBool(key string) bool {
	if !viper.IsSet(key) {
		panic("failed to get environment key: " + key)
	}

	return viper.GetBool(key)
}
