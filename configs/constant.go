package configs

import "github.com/spf13/viper"

var (
	UNIDOC_API_KEY string
)

func init() {
	UNIDOC_API_KEY = viper.GetString("unidoc.api_key")
}
