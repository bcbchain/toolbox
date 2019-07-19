package mydice2winstub

import (
	bcType "blockchain/types"
	"fmt"
	"runtime"
	"strings"

	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/types"
	"contract/stubcommon/common"
	stubType "contract/stubcommon/types"
	tmcommon "github.com/tendermint/tmlibs/common"

	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/rlp"
	"contract/orgexample/code/mydice2win/v1.0/mydice2win"
	"github.com/tendermint/tmlibs/log"
)

//Dice2WinStub an object
type Dice2WinStub struct {
	logger log.Logger
}

var _ stubType.IContractStub = (*Dice2WinStub)(nil)

//New generate a stub
func New(logger log.Logger) stubType.IContractStub {
	return &Dice2WinStub{logger: logger}
}

//FuncRecover recover panic by Assert
func FuncRecover(response *bcType.Response) {
	if err := recover(); err != nil {
		if _, ok := err.(types.Error); ok {
			error := err.(types.Error)
			response.Code = error.ErrorCode
			response.Log = error.Error()
		} else if e, ok := err.(error); ok {
			if strings.HasPrefix(e.Error(), "runtime error") {
				logCaller()
				response.Code = types.ErrStubDefined
				response.Log = e.Error()
			} else {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
}

func logCaller() {
	skip := 0
	pc, callerFile, callerLine, ok := runtime.Caller(skip)
	if !ok {
		return
	}
	var testFile string
	var testLine int
	testFunc := runtime.FuncForPC(pc)
	if runtime.FuncForPC(pc) != testFunc {
		for {
			skip++
			if pc, file, line, ok := runtime.Caller(skip); ok {
				// Note that the test line may be different on
				// distinct calls for the same test.  Showing
				// the "internal" line is helpful when debugging.
				if runtime.FuncForPC(pc) == testFunc {
					testFile, testLine = file, line
					break
				}
			} else {
				break
			}
		}
	}
	if testFile != "" && (testFile != callerFile || testLine != callerLine) {
		fmt.Println(testFile, testLine)
	}
	fmt.Println(callerFile, callerLine)
}

// InitChain initial smart contract
func (pbs *Dice2WinStub) InitChain(smc sdk.ISmartContract) (response bcType.Response) {
	defer FuncRecover(&response)

	contractObj := new(mydice2win.Dice2Win)
	contractObj.SetSdk(smc)
	contractObj.InitChain()

	response.Code = types.CodeOK
	return response
}

// UpdateChain update smart contract
func (pbs *Dice2WinStub) UpdateChain(smc sdk.ISmartContract) (response bcType.Response) {
	defer FuncRecover(&response)
	contractObj := new(mydice2win.Dice2Win)
	contractObj.SetSdk(smc)
	contractObj.UpdateChain()

	response.Code = types.CodeOK
	return response
}

// Mine call mine of smart contract
func (pbs *Dice2WinStub) Mine(smc sdk.ISmartContract) (response bcType.Response) {
	defer FuncRecover(&response)
	contractObj := new(mydice2win.Dice2Win)
	contractObj.SetSdk(smc)
	contractObj.Mine()

	response.Code = types.CodeOK
	return response
}

//Invoke invoke function
func (pbs *Dice2WinStub) Invoke(smc sdk.ISmartContract) (response bcType.Response) {
	return pbs.InvokeInternal(smc, true)
}

//InvokeInterface invoke function
func (pbs *Dice2WinStub) InvokeInternal(smc sdk.ISmartContract, feeFlag bool) (response bcType.Response) {
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
	case "d373a935": // prototype: SetSecretSigner(types.PubKey)
		setSecretSigner(smc)
	case "44f1b25c": // prototype: SetSettings(string)
		setSettings(smc)
	case "b4af57dc": // prototype: SetRecvFeeInfos(string)
		setRecvFeeInfos(smc)
	case "948c4d24": // prototype: WithdrawFunds(string,types.Address,bn.Number)
		withdrawFunds(smc)
	case "d517c92": // prototype: PlaceBet(bn.Number,int64,int64,[]byte,[]byte,types.Address)
		placeBet(smc)
	case "23cf2305": // prototype: SettleBet([]byte)
		settleBet(smc)
	case "de78d51d": // prototype: RefundBet([]byte)
		refundBet(smc)
	default:
		err.ErrorCode = types.ErrInvalidMethod
	}
	response = common.CreateResponse(smc.Message(), response.Tags, data, fee, gasUsed, smc.Tx().GasLimit(), err)
	return
}

func setSecretSigner(smc sdk.ISmartContract) {
	items := smc.Message().Items()
	sdk.Require(len(items) == 1, types.ErrStubDefined, "Invalid message data")
	var err error

	var v0 types.PubKey
	err = rlp.DecodeBytes(items[0], &v0)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	contractObj := new(mydice2win.Dice2Win)
	contractObj.SetSdk(smc)
	contractObj.SetSecretSigner(v0)
}

func setSettings(smc sdk.ISmartContract) {
	items := smc.Message().Items()
	sdk.Require(len(items) == 1, types.ErrStubDefined, "Invalid message data")
	var err error

	var v0 string
	err = rlp.DecodeBytes(items[0], &v0)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	contractObj := new(mydice2win.Dice2Win)
	contractObj.SetSdk(smc)
	contractObj.SetSettings(v0)
}

func setRecvFeeInfos(smc sdk.ISmartContract) {
	items := smc.Message().Items()
	sdk.Require(len(items) == 1, types.ErrStubDefined, "Invalid message data")
	var err error

	var v0 string
	err = rlp.DecodeBytes(items[0], &v0)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	contractObj := new(mydice2win.Dice2Win)
	contractObj.SetSdk(smc)
	contractObj.SetRecvFeeInfos(v0)
}

func withdrawFunds(smc sdk.ISmartContract) {
	items := smc.Message().Items()
	sdk.Require(len(items) == 3, types.ErrStubDefined, "Invalid message data")
	var err error

	var v0 string
	err = rlp.DecodeBytes(items[0], &v0)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	var v1 types.Address
	err = rlp.DecodeBytes(items[1], &v1)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	var v2 bn.Number
	err = rlp.DecodeBytes(items[2], &v2)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	contractObj := new(mydice2win.Dice2Win)
	contractObj.SetSdk(smc)
	contractObj.WithdrawFunds(v0, v1, v2)
}

func placeBet(smc sdk.ISmartContract) {
	items := smc.Message().Items()
	sdk.Require(len(items) == 6, types.ErrStubDefined, "Invalid message data")
	var err error

	var v0 bn.Number
	err = rlp.DecodeBytes(items[0], &v0)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	var v1 int64
	err = rlp.DecodeBytes(items[1], &v1)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	var v2 int64
	err = rlp.DecodeBytes(items[2], &v2)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	var v3 []byte
	err = rlp.DecodeBytes(items[3], &v3)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	var v4 []byte
	err = rlp.DecodeBytes(items[4], &v4)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	var v5 types.Address
	err = rlp.DecodeBytes(items[5], &v5)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	contractObj := new(mydice2win.Dice2Win)
	contractObj.SetSdk(smc)
	contractObj.PlaceBet(v0, v1, v2, v3, v4, v5)
}

func settleBet(smc sdk.ISmartContract) {
	items := smc.Message().Items()
	sdk.Require(len(items) == 1, types.ErrStubDefined, "Invalid message data")
	var err error

	var v0 []byte
	err = rlp.DecodeBytes(items[0], &v0)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	contractObj := new(mydice2win.Dice2Win)
	contractObj.SetSdk(smc)
	contractObj.SettleBet(v0)
}

func refundBet(smc sdk.ISmartContract) {
	items := smc.Message().Items()
	sdk.Require(len(items) == 1, types.ErrStubDefined, "Invalid message data")
	var err error

	var v0 []byte
	err = rlp.DecodeBytes(items[0], &v0)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	contractObj := new(mydice2win.Dice2Win)
	contractObj.SetSdk(smc)
	contractObj.RefundBet(v0)
}
