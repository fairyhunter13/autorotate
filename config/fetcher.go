package config

import (
	"github.com/spf13/viper"
)

const (
	EnvDevicePatternOnly = "device.pattern.only"
)

func GetDevicePatternOnly() string {
	return viper.GetString(EnvDevicePatternOnly)
}
