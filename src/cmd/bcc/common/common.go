package common

import (
	"github.com/pkg/errors"
	"github.com/tendermint/tmlibs/log"
	"os"
)

var (
	bccConfig           BccConfig
	bccRPCServiceConfig BccRPCServiceConfig
	logger              log.Logger
)

func LoadBCCConfig() error {
	configFile := "./.config/bcc.yaml"

	err := InitConfig(&bccConfig, configFile)
	if err != nil {
		return errors.New("Init config fail err info : " + err.Error())
	}

	return nil
}

func InitRPC() error {
	configFile := "./.config/bccRpcService.yaml"
	moduleName := "bccRPCService"

	err := InitConfig(&bccRPCServiceConfig, configFile)
	if err != nil {
		return errors.New("Init config fail err info : " + err.Error())
	}

	initLog(moduleName)

	return nil
}

func GetBCCConfig() BccConfig {
	return bccConfig
}

func GetBCCServiceConfig() BccRPCServiceConfig {
	return bccRPCServiceConfig
}

func initLog(moduleName string) {
	l := log.NewTMLogger("./log", moduleName)
	l.SetOutputToFile(true)
	l.SetOutputToScreen(false)
	l.AllowLevel(bccRPCServiceConfig.LoggerLevel)
	logger = l
}

func GetLogger() log.Logger {
	return logger
}

func OutCertFileIsExist() (string, string) {
	crtPath := "./.config/server.crt"
	keyPath := "./.config/server.key"

	_, err := os.Stat(bccRPCServiceConfig.OutCertPath + ".crt")
	if err != nil {
		return crtPath, keyPath
	}

	_, err = os.Stat(bccRPCServiceConfig.OutCertPath + ".key")
	if err != nil {
		return crtPath, keyPath
	}

	return bccRPCServiceConfig.OutCertPath + ".crt", bccRPCServiceConfig.OutCertPath + ".key"
}
