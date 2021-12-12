package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type LogConfig struct {
	Level    string `yaml:"level"`
	Format   string `yaml:"format"`
	Filename string `yaml:"filename"`
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
	Data_path   string `yaml:"data_path"`
	Cycle_msec  int    `yaml:"cycle_msec"`
	Target_file string `yaml:"target_file"`
}

type CommonConfig struct {
	Log   LogConfig   `yaml:"log"`
	Kafka KafkaConfig `yaml:"kafka"`
}

type ProcConfig struct {
	Process ProcessConfig `yaml:"process"`
}

type Config struct {
	CommonConf CommonConfig
	ProcConf   ProcConfig
}

func fileReader(s_filepath string) []byte {
	bytes, err := ioutil.ReadFile(s_filepath)
	if err != nil {
		fmt.Println("Error Occured: ", err)
		panic(err)
	}

	return bytes
}

func commonConfigReader(filepath string) *CommonConfig {
	var config CommonConfig

	content := fileReader(filepath)
	err := yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("Error : ", err)
	}

	return &config
}

func processConfigReader(filepath string) *ProcConfig {
	var config ProcConfig

	content := fileReader(filepath)
	err := yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("Error : ", err)
	}

	return &config
}

func ConfigReader(filepath string) *Config {
	var config Config
	(config.CommonConf) = *commonConfigReader("../config/common.yaml")
	(config.ProcConf) = *processConfigReader(filepath)

	return &config
}
