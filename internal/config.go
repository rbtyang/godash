package internal

import (
	"github.com/rbtyang/godash/dashfile"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Redis *Redis `yaml:"redis"` //这里Redis首字母要大写，否则读不到配置
}
type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func LoadConfig() (*Config, error) {
	var config = Config{&Redis{}}
	log.Println("Before lock_test.go tests")

	var err error
	var yamlFile *os.File

	yamlFile, err = os.OpenFile("../test.yaml", os.O_RDONLY, os.ModePerm)
	if err != nil {
		yamlFile, err = os.OpenFile("../test.demo.yaml", os.O_RDONLY, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	defer yamlFile.Close()

	yamlFileByt, err := dashfile.ReadByFile(yamlFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFileByt, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
