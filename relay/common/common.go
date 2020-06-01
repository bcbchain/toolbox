package common

import (
	"errors"
	"fmt"
	"github.com/bcbchain/bclib/tendermint/tmlibs/log"
	"github.com/bcbchain/bclib/wal"
	"github.com/bcbchain/toolbox/relay/common/config"
)

var (
	relayConfig config.RelayConfig
	logger      log.Logger

	walName      string
	walPassword  string
	keyStorePath string
)

func InitAll(name, password, keyStore string) error {
	err := checkInputParams(name, password, keyStore)
	if err != nil {
		return err
	}
	walName = name
	walPassword = password
	keyStorePath = keyStore

	moduleName := "relay"
	configPath := "./.config/relay.yaml"
	err = relayConfig.InitConfig(configPath)
	if err != nil {
		return errors.New("Init config fail err info : " + err.Error())
	}

	initLog(moduleName)

	return nil
}

func GetConfig() config.RelayConfig {
	return relayConfig
}

func GetLogger() log.Logger {
	return logger
}

func GetWalletInfo() (string, string, string) {
	return walName, walPassword, keyStorePath
}

func initLog(moduleName string) {
	l := log.NewTMLogger("./log", moduleName)
	l.SetOutputToFile(relayConfig.LoggerFile)
	l.SetOutputToScreen(relayConfig.LoggerScreen)
	l.AllowLevel(relayConfig.LoggerLevel)
	logger = l
}

func checkInputParams(walName, walPassword, keyStorePath string) error {
	if walName == "" {
		return errors.New("Need wallet name ")
	}

	_, err := wal.LoadAccount(keyStorePath, walName, walPassword)
	if err != nil {
		return errors.New(fmt.Sprintf("Load account \"%v\" failed, %v", walName, err.Error()))
	}

	fmt.Println("checkInputParams success!")
	return nil
}

func FuncRecover(errPtr *error) {
	if err := recover(); err != nil {
		msg := ""
		if errInfo, ok := err.(error); ok {
			msg = errInfo.Error()
		}

		if errInfo, ok := err.(string); ok {
			msg = errInfo
		}

		*errPtr = errors.New(msg)
	}
}
