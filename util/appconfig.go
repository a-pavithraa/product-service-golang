package util

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AppSettings struct {
	Port           string         `mapstructure:"port"`
	LogLevel       string         `mapstructure:"logLevel"`
	Regions        []string       `mapstructure:"regions"`
	KeycloakConfig KeycloakConfig `mapstructure:"keycloak"`
	DBConfig       DBConfig       `mapstructure:"db"`
}
type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}
type KeycloakConfig struct {
	Realm          string `mapstructure:"realm"`
	ClientID       string `mapstructure:"clientId"`
	ClientSecret   string `mapstructure:"clientSecret"`
	BaseURL        string `mapstructure:"url"`
	TimeoutSeconds int    `mapstructure:"timeoutSeconds"`
}

func LoadAppConfig() AppSettings {
	var settings AppSettings
	log.Println("Loading Server Configurations...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&settings)
	if err != nil {
		log.Fatal(err)
	}
	return settings

}
