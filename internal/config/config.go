package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

const TTL = time.Hour * 1

type Config struct {
	Env         string `yaml:"env" env-default:"prod"`
	StoragePath string `yaml:"storage_path" env-required:"./data"`
	TokenTTL    time.Duration
	GRPC        GRPConfig `yaml:"grpc"`
}

type GRPConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {

	path := "/root/apps/grpc-auth/config/prod.yaml"

	return MustLoadByPath(path)
}

func MustLoadByPath(configPath string) *Config {
	fmt.Println(configPath)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exists" + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("failed to read config" + err.Error())
	}
	cfg.SetTTL()

	return &cfg
}

func (c *Config) SetTTL() {
	c.TokenTTL = TTL

	if c.TokenTTL == 0 {
		panic("TTL is not set")
	}
}
