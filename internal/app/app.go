package app

import (
	"net/http"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	app_http "github.com/Vaelatern/gokrazy-statuspage/internal/http"
)

func initConfig() error {
	viper.SetDefault("listen", ":3000")
	viper.SetDefault("base-url", "/")
	viper.SetDefault("poll-frequency", 30)
	viper.BindPFlags(pflag.CommandLine)
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	viper.WatchConfig()
	return nil
}

func Entrypoint() error {
	if err := initConfig(); err != nil {
		return err
	}
	return http.ListenAndServe(viper.GetString("listen"), app_http.Router(viper.GetString("base-url")))
}
