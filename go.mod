module github.com/bcbchain/toolbox

go 1.14

require (
	github.com/bcbchain/bcbchain v0.0.0-20200602064057-474fc36dc199
	github.com/bcbchain/bclib v0.0.0-20200604101126-5bd6cff2be45
	github.com/bcbchain/sdk v0.0.0-20200602073908-61a228ca7831
	github.com/bcbchain/tendermint v0.0.0-20200602064030-2579a5bffc23
	github.com/btcsuite/btcutil v1.0.2
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.0.0
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/bcbchain/bcbchain => ../bcbchain
