package hue

import "github.com/spf13/viper"

var (
	username string
	url      string
)

func List() string {

	username = viper.GetString("hue.username")
	url = viper.GetString("hue.url")
	return url + "/" + username
}
