package lib

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type LogConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

type KafkaConfig struct {
	Hosts   []string `yaml:"hosts"`
	Port    int      `yaml:"port"`
	Timeout int      `yaml:"timeout"`
}

type ProcessConfig struct {
	Name        string `yaml:"name"`
	Topic_name  string `yaml:"topic_name"`
	Script_path string `yaml:"script_path"`
}

type Config struct {
	Log     LogConfig     `yaml:"log"`
	Kafka   KafkaConfig   `yaml:"kafka"`
	Process ProcessConfig `yaml:"process"`
}

func FileReader(s_filepath string) []byte {
	bytes, err := ioutil.ReadFile(s_filepath)
	if err != nil {
		fmt.Println("Error Occured: ", err)
		panic(err)
	}

	return bytes
}

func ConfigReader(filepath string) *Config {
	var config Config

	content := FileReader(filepath)
	err := yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("Error : ", err)
	}

	return &config
}
