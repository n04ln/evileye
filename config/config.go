package config

import "github.com/kelseyhightower/envconfig"

var (
	CommitHash string
	BuildAt    string
)

type Config struct {
	Salt      string `default:"chu2byo"`
	Secret    string `default:"welcome_to_underground"`
	Threshold int64  `default:10`
}

var config Config

func init() {
	config, _ = Init()
}

func Init() (Config, error) {
	err := envconfig.Process("Evileye", &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func GetConfig() Config {
	return config
}
