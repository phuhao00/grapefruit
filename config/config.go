package config

import (
	"context"
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"io"
	"os"
)

type Config struct {
	Redis      *Redis
	Mongo      *Mongo
	Es         *ES
	Postgresql *Postgresql
}

var config *Config

func GetFileName(ctx context.Context) string {
	env := GetEnv(ctx)
	fileName := "dev.toml"
	if env == "dev" {
		fileName = "dev.toml"
	}
	if env == "prod" {
		fileName = "prod.toml"
	}
	if env == "qa" {
		fileName = "prod.toml"
	}
	return "./config/" + fileName
}

func LoadConfig(ctx context.Context) (*Config, error) {
	var c Config
	filename := GetFileName(ctx)
	configFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := configFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	bytes, err := io.ReadAll(configFile)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(bytes, &c)
	if err != nil {
		return &c, err
	}
	config = &c
	return config, nil
}

func GetEnv(ctx context.Context) string {
	env := os.Getenv("ENV")
	return env
}

func GormGenerateLoadConfig(filename string) (*Config, error) {
	var c Config
	configFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := configFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	bytes, err := io.ReadAll(configFile)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(bytes, &c)
	if err != nil {
		return &c, err
	}
	config = &c
	return config, nil
}
