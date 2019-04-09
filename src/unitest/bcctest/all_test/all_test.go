package all

import (
	"cmd/bcc/common"
	"testing"

	"math/rand"
	"time"

	"gopkg.in/check.v1"
)

var (
	// 合法参数定义
	name                  = "qingzhong"
	password              = "qingzhong361001G@"
	ownerName             = "owner"
	ownerPassword         = "Ab1@Cd3$"
	gasLimit              = "10000000"
	note                  = ""
	chainID               = ""
	keystorePath          = ".keystore"
	token                 = "LOC"
	to                    = "localAoMsg7xiJ63zBxkjWEfyBh6i7mVL5QmCK"
	value                 = "1000000"
	contractName          = "mybasictype"
	pubKey                = "[\"0x01bd6c29d63f5f32aa33955f26a28459988edea4de517f77372e77db33958e6e\"]"
	PrivateKey            = "832c3477b3a730e9601ed774a88c27ca43112ff7d0718686a33c21721dd3369701bd6c29d63f5f32aa33955f26a28459988edea4de517f77372e77db33958e6e"
	deployer              = "localL9BzYNYns5VCRaJgfHEBJLzS1bhpHjx7j"
	orgName               = "example"
	version               = "1.0"
	codeFile              = "./mydice2win-1.0.tar.gz"
	effectHeight          = "1000"
	owner                 = "localL9BzYNYns5VCRaJgfHEBJLzS1bhpHjx7j"
	orgID                 = "orgBtjfCSPCAJ84uQWcpNr74NLMWYm5SXzer"
	contractAddr          = "localALG8jX5rpSVTYUjSwxQFrrAegFEMoFwZz"
	totalSupply           = "100000000000000"
	gasprice              = "2500"
	addSupplyEnabled      = "true"
	burnEnabled           = "true"
	txHash                = "3f4e72dc0485af3a31c1291340ea37740e617389f32a14c717e5b8664deeb344"
	height                = "2"
	accAddr               = "localL9BzYNYns5VCRaJgfHEBJLzS1bhpHjx7j"
	hexStr                = "abcdefghijlkmnopqrstuvwxyz9876543210"
	all                   = "false"
	contractNamedemo1     = "mybasictype"
	contractNamedemo1Addr = "localALG8jX5rpSVTYUjSwxQFrrAegFEMoFwZz"
	contractNamedemo2     = "mymixtype"

	// 工具定义
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func init() {
	err := common.LoadBCCConfig()
	if err != nil {
		panic(err)
	}
	err = common.InitRPC()
	if err != nil {
		panic(err)
	}

}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { check.TestingT(t) }

type MySuite struct{}

var _ = check.Suite(&MySuite{})
