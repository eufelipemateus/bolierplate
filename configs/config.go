package config

import (
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API           aPiConfig
	DB            dBConfig
	SenGridConfig sengridConfig
	Recaptcha     recaptcha
}

type aPiConfig struct {
	Port string
	Name string
	GenerateDB bool
}

type dBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

type sengridConfig struct {
	SengridApiKey string
}

type recaptcha struct {
	Secret  string
	SiteKey string
}

func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "locahost")
	viper.SetDefault("database.port", "5432")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)
	cfg.API = aPiConfig{
		Port: viper.GetString("api.port"),
		Name: viper.GetString("api.name"),
		GenerateDB: viper.GetBool("api.generateDB"),
	}

	cfg.DB = dBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}

	cfg.SenGridConfig = sengridConfig{
		SengridApiKey: viper.GetString("sengrid.apiKey"),
	}

	cfg.Recaptcha = recaptcha{
		Secret:  viper.GetString("recaptcha.secret"),
		SiteKey: viper.GetString("recaptcha.sitekey"),
	}


	return nil
}

func GetDB() dBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}

func GetAppName() string {
	return cfg.API.Name
}

func IsGenerateDB() bool {
	return cfg.API.GenerateDB
}

func GetSengridApiKey() string {
	return cfg.SenGridConfig.SengridApiKey
}

func GetRecaptcha() recaptcha {
	return cfg.Recaptcha
}
