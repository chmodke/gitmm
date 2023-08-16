package config

import (
	"fmt"
	"github.com/chmodke/gitmm/log"
	"github.com/spf13/viper"
	"os"
)

var viperConfig *viper.Viper

var (
	Remote   map[string]string
	Origin   string
	Upstream string
	Repos    []string
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
		log.Consoleln(err)
		log.Consoleln("可以执行`gitmm config`命令生成示例配置文件。")
		os.Exit(1)
	}
	Remote = viperConfig.GetStringMapString("remote")
	var ok = false
	if Origin, ok = Remote["origin"]; !ok {
		log.Consoleln("未配置origin远端地址")
		os.Exit(1)
	}
	if Upstream, ok = Remote["upstream"]; !ok {
		log.Consoleln("未配置upstream远端地址")
		os.Exit(1)
	}
	Repos = viperConfig.GetStringSlice("repos")
}

func WriteCfg() {
	sample := viper.New()
	remote := map[string]string{"upstream": "git@gitee.com:chmodke", "origin": "git@github.com:chmodke"}
	sample.Set("remote", remote)
	sample.Set("repos", []string{"arpc", "ftrans"})

	sample.WriteConfigAs("repo_sample.yaml")
}
