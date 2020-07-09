package log

import (
	"github.com/bcbchain/bclib/tendermint/tmlibs/log"
	"github.com/bcbchain/toolbox/bcscan/conf"
)

var (
	logger log.Loggerf
)

func InitLogger(c *conf.BcScanConfig) {
	if c == nil {
		panic("config can not be nil")
	}

	logger = log.NewTMLogger("./log", "bcscan")
	logger.SetOutputToScreen(false)
	logger.SetOutputToFile(c.LoggerFile)
	logger.AllowLevel(c.LoggerLevel)
}

func GetLogger() log.Loggerf {
	return logger
}
