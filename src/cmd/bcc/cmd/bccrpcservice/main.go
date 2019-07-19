package main

import (
	"cmd/bcc/cache"
	"cmd/bcc/common"
	"cmd/bcc/core"
	"common/rpc/lib/server"
	"github.com/tendermint/go-amino"
	cmn "github.com/tendermint/tmlibs/common"
	"net/http"
	"os"
)

func main() {
	err := common.LoadBCCConfig()
	if err != nil {
		panic(err)
	}

	err = common.InitRPC()
	if err != nil {
		panic(err)
	}

	cache.Init(".keystore")

	rpcLogger := common.GetLogger()

	coreCodec := amino.NewCodec()

	mux := http.NewServeMux()

	rpcserver.RegisterRPCFuncs(mux, routes, coreCodec, rpcLogger)

	if common.GetBCCServiceConfig().UseHttps {
		crtPath, keyPath := common.OutCertFileIsExist()
		_, err = rpcserver.StartHTTPAndTLSServer(common.GetBCCServiceConfig().ServerAddr, mux, crtPath, keyPath, rpcLogger)
		if err != nil {
			cmn.Exit(err.Error())
		}
	} else {
		_, err = rpcserver.StartHTTPServer(common.GetBCCServiceConfig().ServerAddr, mux, rpcLogger)
		if err != nil {
			cmn.Exit(err.Error())
		}
	}

	// Wait forever
	cmn.TrapSignal(func(signal os.Signal) {
	})
}

var routes = map[string]*rpcserver.RPCFunc{
	// bcc api
	"bcc_registerOrg":    rpcserver.NewRPCFunc(core.RegisterOrg, "name,password,bccParams"),
	"bcc_setOrgSigners":  rpcserver.NewRPCFunc(core.SetOrgSigners, "name,password,bccParams"),
	"bcc_setOrgDeployer": rpcserver.NewRPCFunc(core.SetOrgDeployer, "name,password,bccParams"),
	"bcc_deployContract": rpcserver.NewRPCFunc(core.DeployContract, "name,password,bccParams"),
	"bcc_registerToken":  rpcserver.NewRPCFunc(core.RegisterToken, "name,password,bccParams"),
	"bcc_transfer":       rpcserver.NewRPCFunc(core.Transfer, "name,password,bccParams"),
	"bcc_call":           rpcserver.NewRPCFunc(core.Call, "name,password,bccParams"),

	// block chain query api
	"bcc_query":        rpcserver.NewRPCFunc(core.Query, "path,data,height,trusted,chainID"),
	"bcc_queryWithKey": rpcserver.NewRPCFunc(core.QueryOfRpc, "key,chainID"),
	"bcc_blockHeight":  rpcserver.NewRPCFunc(core.BlockHeight, "chainID"),
	"bcc_block":        rpcserver.NewRPCFunc(core.BlockForRpc, "height,bTime,num,chainID"),
	"bcc_transaction":  rpcserver.NewRPCFunc(core.Transaction, "txHash,chainID"),
	"bcc_balance":      rpcserver.NewRPCFunc(core.Balance, "address,name,password,tokenName,all,chainID,keyStorePath"),
	"bcc_nonce":        rpcserver.NewRPCFunc(core.Nonce, "address,name,password,chainID,keyStorePath"),
	"bcc_commitTx":     rpcserver.NewRPCFunc(core.CommitTx, "tx,chainID"),
	"bcc_version":      rpcserver.NewRPCFunc(core.Version, ""),
	"bcc_contractInfo": rpcserver.NewRPCFunc(core.ContractInfoForRPC, "orgName,contractName,orgID,contractAddr,chainID"),
}
