package configs

import (
	"github.com/spf13/viper"
	"log"
)

type config struct {
	WeatherApiKey       string `mapstructure:"WEATHER_API_KEY"`
	WeatherApiUrl       string `mapstructure:"WEATHER_API_URL"`
	CepApiUrl           string `mapstructure:"CEP_API_URL"`
	LimiterReqPerSecIP  int    `mapstructure:"LIMITER_REQ_PER_SEC_IP"`
	LimiterReqPerSecKey int    `mapstructure:"LIMITER_REQ_PER_SEC_KEY"`
	LimiterPenaltySec   int    `mapstructure:"LIMITER_PENALTY_SEC"`
	CacheDbUrl          string `mapstructure:"CACHE_DB_URL"`
	CacheDbPassword     string `mapstructure:"CACHE_DB_PASSWORD"`
}

var (
	Configs *config
)

func init() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("../..")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	e := viper.ReadInConfig()
	if e != nil {
		log.Fatal("Can't find the file app.env : ", e)
	}

	e = viper.Unmarshal(&Configs)
	if e != nil {
		log.Fatal("Can't unmarshal the file app.env : ", e)
	}
}
