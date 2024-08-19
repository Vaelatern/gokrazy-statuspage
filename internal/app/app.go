package app

import (
	"fmt"
	"net/http"
	"os"

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

	if otherConfigDir, isSet := os.LookupEnv("CONFIG_DIR"); isSet {
		viper.AddConfigPath(otherConfigDir)
	}
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	viper.WatchConfig()
	return nil
}

func Entrypoint() error {
	if err := initConfig(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Listening on ", viper.GetString("listen"), " and at base URL ", viper.GetString("base-url"))
	return http.ListenAndServe(viper.GetString("listen"), app_http.Router(viper.GetString("base-url")))
}
