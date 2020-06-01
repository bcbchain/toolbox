package relay

import "github.com/bcbchain/toolbox/relay/common"

// 启动
func Start() {
	common.GetLogger().Debug("relay starting...")

	StartScanner()

	StartCarrier()
}
