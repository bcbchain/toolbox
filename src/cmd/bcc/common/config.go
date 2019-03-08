package common

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type BccConfig struct {
	DefaultChainID string   `yaml:"defaultChainID"`
	Bcb            []string `yaml:"bcb"`
	Bcbtest        []string `yaml:"bcbtest"`
	Devtest        []string `yaml:"devtest"`
	Local          []string `yaml:"local"`
}

type BccRPCServiceConfig struct {
	ServerAddr  string `yaml:"serverAddr"`
	UseHttps    bool   `yaml:"useHttps"`
	OutCertPath string `yaml:"outCertPath"`
	LoggerLevel string `yaml:"loggerLevel"`
}

func InitConfig(c interface{}, configFile string) error {
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("yamlFile.Get err #%v\n ", err)
		return err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Printf("Unmarshal: %v\n", err)
		return err
	}

	return nil
}
