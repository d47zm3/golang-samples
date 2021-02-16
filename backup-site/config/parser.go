package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Remote struct {
		Ssh struct {
			Hostname string `yaml:"hostname"`
			Username string `yaml:"username"`
			KeyPath  string `yaml:"keypath"`
		} `yaml:"ssh"`
		Backup struct {
			ScriptPath  string `yaml:"script_path"`
			Destination string `yaml:"destination"`
			Prefix      string `yaml:"prefix"`
		} `yaml:"backup"`
	} `yaml:"remote"`
	Local struct {
		BackupDestination string `yaml:"backup_destination"`
	} `yaml:"local"`
	AwsBucketName string `yaml:"aws_bucket_name"`
	SlackWebhook  string `yaml:"slack_webhook"`
}

func ParseConfig(path string) Config {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Printf(err.Error())
	}

	return cfg
}
