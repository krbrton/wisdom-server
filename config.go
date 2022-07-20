package wisdom_server

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

const (
	DefaultConfigPath = ".wisdom-server"
	DefaultConfigFile = "config.yaml"
)

type Config struct {
	Host      string `yaml:"host"`
	Port      uint32 `yaml:"port"`
	Timeout   int64  `yaml:"timeout"`
	SecretKey string `yaml:"secret_key"`
}

func ReadConfig() (*Config, error) {
	config := &Config{}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configPath := path.Join(homeDir, DefaultConfigPath, DefaultConfigFile)

	if err = cleanenv.ReadConfig(configPath, config); err != nil {
		return nil, err
	}

	logrus.WithField("path", configPath).Info("Config loaded")

	return config, nil
}
