package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type SettingsSchema struct {
	Key    string `yaml:"api_key"`
	Limits struct {
		PerSecond int `yaml:"per_second"`
	} `yaml:"limits"`
}

var settings SettingsSchema

func init() {
	content, err := ioutil.ReadFile("./settings.yml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(content, &settings)
	if err != nil {
		log.Fatal(err)
	}
}

func GetKey() string {
	return settings.Key
}
