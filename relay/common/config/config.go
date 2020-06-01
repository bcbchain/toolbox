package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type RelayConfig struct {
	MainChainID  string   `yaml:"mainChainID"`
	MainChainUrl []string `yaml:"mainChainUrl"`
	SideChainIDs []string `yaml:"sideChainIDs"`
	DBName       string   `yaml:"dbName"`

	LoggerScreen bool   `yaml:"loggerScreen"`
	LoggerFile   bool   `yaml:"loggerFile"`
	LoggerLevel  string `yaml:"loggerLevel"`
}

func (c *RelayConfig) InitConfig(configFile string) error {
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
