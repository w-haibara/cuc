package config

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var Dir = ConfigDir()

func ConfigDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}

	dir := filepath.Join(home, ".config", "cuc")
	if _, err := os.Stat(dir); err != nil {
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			panic(err.Error())
		}
	}

	return dir
}

func ReadAPIKey() string {
	type tml struct {
		Key string
	}
	t := new(tml)
	if _, err := toml.DecodeFile(filepath.Join(Dir, "key.toml"), t); err != nil {
		panic(err.Error())
	}
	return t.Key
}

type Config struct {
	Team   string
	Space  string
	Splint SplintConfig
}

type SplintConfig struct {
	Folder     string
	TimeFormat string `toml:"time_format"`
}

func ReadConfig() Config {
	config := new(Config)
	if _, err := toml.DecodeFile(filepath.Join(Dir, "config.toml"), config); err != nil {
		panic(err.Error())
	}
	return *config
}
