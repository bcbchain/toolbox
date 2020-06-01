package conf

import (
	"errors"
	"fmt"
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v2"
)

type BcScanConfig struct {
	ChainID                   string   `yaml:"chainID"`
	URLs                      []string `yaml:"urls"`
	ConcurrentNumberOfOneNode int      `yaml:"concurrentNumberOfOneNode"`
	DBPort                    string   `yaml:"dbPort"`

	LoggerFile  bool   `yaml:"loggerFile"`
	LoggerLevel string `yaml:"loggerLevel"`
}

var (
	BCScanConf *BcScanConfig
	once       sync.Once
)

func LoadBCSConfig() error {
	configFile := "./.config/bcscan.yaml"
	var err error
	once.Do(func() {
		BCScanConf = new(BcScanConfig)
		err = initConfig(BCScanConf, configFile)
	})
	if err != nil {
		return err
	}

	if BCScanConf.ChainID == "" {
		return errors.New("config chainID can mot be empty")
	}

	if len(BCScanConf.URLs) == 0 {
		return errors.New("config URLs can not be empty")
	}

	if BCScanConf.ConcurrentNumberOfOneNode <= 0 || BCScanConf.ConcurrentNumberOfOneNode > 10 {
		//return errors.New("ConcurrentNumberOfOneNode must be between 0 and 10")
	}

	return nil
}

func GetConfig() *BcScanConfig {
	return BCScanConf
}

func initConfig(c interface{}, configFile string) error {
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
