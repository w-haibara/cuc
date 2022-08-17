package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var (
	configFilePath = DefaultCondigFilePath()
	defaultConfigs = []Config{
		{
			Key:           "key",
			Description:   "description",
			DefaultValue:  "default value",
			AllowedValues: []string{"default value", "val1", "val2"},
		},
	}
	configs          = []Config{}
	currentConfigMap sync.Map
)

func Init() {
	if _, err := os.Stat(configFilePath); err != nil {
		if err := os.MkdirAll(filepath.Dir(configFilePath), 0755); err != nil {
			panic(fmt.Sprint("faild to create user config dir: ", err.Error()))
		}

		b, err := json.MarshalIndent(defaultConfigs, "", "  ")
		if err != nil {
			panic(fmt.Sprint("faild to unmarshal default config json: ", err.Error()))
		}

		if err := os.WriteFile(configFilePath, b, 0666); err != nil {
			panic(fmt.Sprint("faild to initilize config file: ", err.Error()))
		}
	}

	b, err := os.ReadFile(configFilePath)
	if err != nil {
		panic(fmt.Sprint("faild to read config file: ", err.Error()))
	}

	if err := json.Unmarshal(b, &configs); err != nil {
		panic(fmt.Sprint("faild to unmarshal config json: ", err.Error()))
	}

	for _, config := range configs {
		val := config.DefaultValue
		if config.Value != nil {
			val = *config.Value
		}

		Set(config.Key, val)
	}
}

func DefaultCondigFilePath() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		panic(fmt.Sprint("faild to get user config dir: ", err.Error()))
	}

	return filepath.Join(dir, "cuc", "config.json")
}

func Get(key string) (string, error) {
	val, ok := currentConfigMap.Load(key)
	if !ok {
		return "", errors.New("faild to get value")
	}

	str, ok := val.(string)
	if !ok {
		return "", errors.New("invalid value")
	}

	return str, nil
}

func Set(key, value string) {
	currentConfigMap.Store(key, value)
}

func Write() error {
	oldConfigs := Configs()
	configs := make([]Config, len(oldConfigs))
	for i, config := range oldConfigs {
		v, err := Get(config.Key)
		if err != nil {
			return err
		}

		config.Value = &v
		configs[i] = config
	}

	b, err := json.MarshalIndent(configs, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(configFilePath, b, os.ModePerm); err != nil {
		return err
	}

	return nil
}

type Config struct {
	Key           string
	Description   string
	Value         *string
	DefaultValue  string
	AllowedValues []string
}

func Configs() []Config {
	return configs
}
