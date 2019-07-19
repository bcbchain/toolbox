package myballotstub

import (
	bcType "blockchain/types"

	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/types"
	"contract/stubcommon/common"
	stubType "contract/stubcommon/types"
	tmcommon "github.com/tendermint/tmlibs/common"

	"blockchain/smcsdk/sdk/jsoniter"
	"blockchain/smcsdk/sdk/rlp"
	"contract/orgexample/code/myballot/v1.0/myballot"
	"github.com/tendermint/tmlibs/log"
)

//BallotStub an object
type BallotStub struct {
	logger log.Logger
}

var _ stubType.IContractStub = (*BallotStub)(nil)

//New generate a stub
func New(logger log.Logger) stubType.IContractStub {
	return &BallotStub{logger: logger}
}

//FuncRecover recover panic by Assert
func FuncRecover(response *bcType.Response) {
	if err := recover(); err != nil {
		if _, ok := err.(types.Error); ok {
			error := err.(types.Error)
			response.Code = error.ErrorCode
			response.Log = error.Error()
		} else {
			panic(err)
		}
	}
}

// InitChain initial smart contract
func (pbs *BallotStub) InitChain(smc sdk.ISmartContract) (response bcType.Response) {
	defer FuncRecover(&response)

	contractObj := new(myballot.Ballot)
	contractObj.SetSdk(smc)
	contractObj.InitChain()

	response.Code = types.CodeOK
	return response
}

// UpdateChain update smart contract
func (pbs *BallotStub) UpdateChain(smc sdk.ISmartContract) (response bcType.Response) {
	defer FuncRecover(&response)

	response.Code = types.CodeOK
	return response
}

// UpdateChain update smart contract
func (pbs *BallotStub) Mine(smc sdk.ISmartContract) (response bcType.Response) {
	defer FuncRecover(&response)

	response.Code = types.CodeOK
	return response
}

//Invoke invoke function
func (pbs *BallotStub) Invoke(smc sdk.ISmartContract) (response bcType.Response) {
	return pbs.InvokeInternal(smc, true)
}

//Invoke invoke function
func (pbs *BallotStub) InvokeInternal(smc sdk.ISmartContract, feeFlag bool) (response bcType.Response) {
	defer FuncRecover(&response)

	// 生成手续费收据
	fee, gasUsed, feeReceipt, err := common.FeeAndReceipt(smc, feeFlag)
	response.Fee = fee
	response.GasUsed = gasUsed
	response.Tags = append(response.Tags, tmcommon.KVPair{Key: feeReceipt.Key, Value: feeReceipt.Value})
	if err.ErrorCode != types.CodeOK {
		response = common.CreateResponse(smc.Message(), response.Tags, "", fee, gasUsed, smc.Tx().GasLimit(), err)
		return
	}

	var data string
	err = types.Error{ErrorCode: types.CodeOK}
	switch smc.Message().MethodID() {
	case "8d31795c": // prototype: Init([]string)
		Init(smc)
	case "b319ae5e": // prototype: GiveRightToVote(types.Address)
		giveRightToVote(smc)
	case "b23ceff6": // prototype: Delegate(types.Address)
		delegate(smc)
	case "b46773ff": // prototype: Vote(uint)
		vote(smc)
	case "1c687fb1": // prototype: WinningProposal()
		data = winningProposal(smc)
	case "de881b9": // prototype: WinnerName()
		data = winnerName(smc)
	default:
		err.ErrorCode = types.ErrInvalidMethod
	}
	response = common.CreateResponse(smc.Message(), response.Tags, data, fee, gasUsed, smc.Tx().GasLimit(), err)
	return
}

func Init(smc sdk.ISmartContract) {
	items := smc.Message().Items()
	sdk.Require(len(items) == 1, types.ErrStubDefined, "Invalid message data")
	var err error

	var v0 []string
	err = rlp.DecodeBytes(items[0], &v0)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	contractObj := new(myballot.Ballot)
	contractObj.SetSdk(smc)
	contractObj.Init(v0)
}

func giveRightToVote(smc sdk.ISmartContract) {
	items := smc.Message().Items()
	sdk.Require(len(items) == 1, types.ErrStubDefined, "Invalid message data")
	var err error

	var v0 types.Address
	err = rlp.DecodeBytes(items[0], &v0)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	contractObj := new(myballot.Ballot)
	contractObj.SetSdk(smc)
	contractObj.GiveRightToVote(v0)
}

func delegate(smc sdk.ISmartContract) {
	items := smc.Message().Items()
	sdk.Require(len(items) == 1, types.ErrStubDefined, "Invalid message data")
	var err error

	var v0 types.Address
	err = rlp.DecodeBytes(items[0], &v0)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	contractObj := new(myballot.Ballot)
	contractObj.SetSdk(smc)
	contractObj.Delegate(v0)
}

func vote(smc sdk.ISmartContract) {
	items := smc.Message().Items()
	sdk.Require(len(items) == 1, types.ErrStubDefined, "Invalid message data")
	var err error

	var v0 uint
	err = rlp.DecodeBytes(items[0], &v0)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	contractObj := new(myballot.Ballot)
	contractObj.SetSdk(smc)
	contractObj.Vote(v0)
}

func winningProposal(smc sdk.ISmartContract) string {
	items := smc.Message().Items()
	sdk.Require(len(items) == 0, types.ErrStubDefined, "Invalid message data")

	contractObj := new(myballot.Ballot)
	contractObj.SetSdk(smc)
	rst0 := contractObj.WinningProposal()
	resultList := make([]interface{}, 0)
	resultList = append(resultList, rst0)

	resBytes, _ := jsoniter.Marshal(resultList)
	return string(resBytes)
}

func winnerName(smc sdk.ISmartContract) string {
	items := smc.Message().Items()
	sdk.Require(len(items) == 0, types.ErrStubDefined, "Invalid message data")

	contractObj := new(myballot.Ballot)
	contractObj.SetSdk(smc)
	rst0 := contractObj.WinnerName()
	resultList := make([]interface{}, 0)
	resultList = append(resultList, rst0)

	resBytes, _ := jsoniter.Marshal(resultList)
	return string(resBytes)
}
