package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var viperConfig *viper.Viper

var (
	OriginGroup string
	MainGroup   string
	Repos       []string
)

func Load(configName string) (*viper.Viper, error) {
	if configName == "" {
		return nil, fmt.Errorf("config name is empty")
	}
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("config file %s.%s not found", configName, "yaml")
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
		fmt.Println("可以执行`gitmm config`命令生成示例配置文件。")
		os.Exit(1)
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
