package cmd

import (
	"github.com/chuck-danner/hue/hue"
	"github.com/spf13/viper"
)

func DefaultHue() *hue.Hue {
	return &hue.Hue{
		Key: viper.GetString("hue.key"),
		URL: viper.GetString("hue.url")}
}
