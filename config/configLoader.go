package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var viperConfig *viper.Viper

var (
	OriginGroup string
	MainGroup   string
	Repos       []string
)

func Load(configName string) (*viper.Viper, error) {
	if configName == "" {
		return nil, fmt.Errorf("config directory is empty")
	}
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("config file %s not found", configName)
		} else {
			return nil, err
		}
	}

	viperConfig = viper.GetViper()
	return viperConfig, nil
}

func LoadCfg() {
	var err error
	viperConfig, err = Load("repo")
	if err != nil {
		fmt.Println(err)
		fmt.Println("please run gitmm config ")
		return
	}
	OriginGroup = viperConfig.GetString("origin_group")
	MainGroup = viperConfig.GetString("main_group")
	Repos = viperConfig.GetStringSlice("repos")
}

func WriteCfg() {
	sample := viper.New()
	sample.Set("main_group", "git@gitee.com:chmodke")
	sample.Set("origin_group", "ssh://git@192.168.100.100:2222/chmodke")
	sample.Set("repos", []string{"arpc", "ftrans"})

	sample.WriteConfigAs("repo_sample.yaml")
}
