package main

import (
	"errors"
	"github.com/bcbchain/bclib/tendermint/tmlibs/log"
	"github.com/bcbchain/toolbox/bbm/rpcclient"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type BbmConfig struct {
	NodeAddr          string `yaml:"nodeAddr"`
	LoggerScreen      bool   `yaml:"loggerScreen"`
	LoggerFile        bool   `yaml:"loggerFile"`
	LoggerFilePath    string `yaml:"loggerFilePath"`
	LoggerLevel       string `yaml:"loggerLevel"`
	OrgSigWallet      string `yaml:"orgSigWallet"`
	OrgSigPassword    string `yaml:"orgSigPassword"`
	GenesisKey        string `yaml:"genesisKey"`
	PrintTPSFrequency int64  `yaml:"printTPSFrequency"`
	txVersion         int64
}

var (
	// bbm config
	bbmConfig BbmConfig
	logger    log.Loggerf
)

func LoadBBMConfig(filePath string) error {
	err := InitConfig(&bbmConfig, filePath)
	if err != nil {
		return errors.New("Init config fail err info : " + err.Error())
	}

	// use 10 for default frequency
	if bbmConfig.PrintTPSFrequency == 0 {
		bbmConfig.PrintTPSFrequency = 10
	}

	// set default private key
	if bbmConfig.GenesisKey == "" {
		bbmConfig.GenesisKey = "012e5aebe1726781ecb702771f9370946356d9e1813bffa1ae7b6c05a800f406704ac09f62376e340212250fb1fb7dae4c442725fbf4277310bbb3b5d8e64ff5"
	}

	return nil
}

func InitConfig(c interface{}, configFile string) error {
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return err
	}

	return nil
}

func InitTXVersion() error {
	client := rpcclient.InitClient(bbmConfig.NodeAddr, logger)

	b := client.Health()
	if 0 == b.ChainVersion {
		logger.Debug("Get BcbChain Version error, use version 1.0")
		bbmConfig.txVersion = 1
	} else {
		bbmConfig.txVersion = b.ChainVersion
	}
	//logger.Info("Get Node Info", "ChainCode Version", bbmConfig.txVersion)
	return nil
}
