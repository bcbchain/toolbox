package common

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type BccConfig struct {
	DefaultChainID string              `yaml:"defaultChainID"`
	Urls           map[string][]string `yaml:"urls"`
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
