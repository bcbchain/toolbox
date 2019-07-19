package testsdk

import (
	"math/rand"
	"testing"

	"blockchain/smcsdk/sdk"
	. "blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
	ut "blockchain/smcsdk/utest"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

var orgID = "orgAJrbk6Wdf7TCbunrXXS5kKvbWVszhC1T"

func (mysuit *MySuite) TestTestSdk_ContractData(c *C) {
	ut.Init(orgID)
	contractOwner := ut.DeployContract(c, contractName, orgID, contractMethods, contractMethods)
	test := NewTestObject(contractOwner)

	testContractInfo(test, contractOwner)
	//
	testContractSetting(test, contractOwner)
}

func testContractInfo(test *TestObject, owner sdk.IAccount) {
	//Address
	//Account
	//Checking contract Owner
	ut.AssertEquals(test.obj.sdk.Message().Contract().Owner().Address(), owner.Address())
	//Checking contract name
	ut.AssertEquals(test.obj.sdk.Message().Contract().Name(), contractName)
	ut.AssertEquals(test.obj.sdk.Message().Contract().KeyPrefix(), "/"+contractName)
	//Version
	//EffectHeight
	//LoseHeight
	//Methods
	//Token
}

//testContractSetting test SetOwner function
func testContractSetting(test *TestObject, owner sdk.IAccount) {
	userWilliam := test.obj.sdk.Helper().BlockChainHelper().CalcAccountFromName("william", "")
	userJack := test.obj.sdk.Helper().BlockChainHelper().CalcAccountFromName("jack", "")

	accWilliam := test.obj.sdk.Helper().AccountHelper().AccountOf(userWilliam)
	accJack := test.obj.sdk.Helper().AccountHelper().AccountOf(userJack)

	var cases = []struct {
		sender   sdk.IAccount
		newOwner types.Address
		dbOwner  types.Address
		excepted uint32
	}{
		{accWilliam, accJack.Address(), owner.Address(), types.ErrNoAuthorization},
		{owner, accWilliam.Address(), accWilliam.Address(), types.CodeOK},
		{accWilliam, owner.Address(), owner.Address(), types.CodeOK},
	}

	for _, c := range cases {
		ut.AssertError(test.run().setSender(c.sender).SetOwner(c.newOwner), c.excepted)
		ut.AssertEquals(test.obj.sdk.Message().Contract().Owner().Address(), c.dbOwner)
	}
}

//TODO: 单元测试中没有构造tx 和 message，故该部分接口在合约运行服务的单元测试中测试

// 固定测试用例，生成一定数量的区块，测试block功能接口
func setMystruct(test *TestObject, sender sdk.IAccount) {

	_pr := test.obj.sdk.Message().Contract().KeyPrefix()
	var cases = []struct {
		value string
	}{
		{"12345"},
		{"23456"},
		{"34567"},
		{"45678"},
	}
	md := Mystruct{}
	md.Mymap = make(map[int64]string)
	for i, c := range cases {
		md.Index = int64(i + 1)
		md.Data = c.value
		md.Mymap[md.Index] = md.Data

		ut.AssertOK(test.run().setSender(sender).SetMydata(md))
		ut.AssertSDB(_pr+"/mydata", md)
		v, _ := test.run().setSender(sender).GetMydata()
		ut.AssertEquals(v.Data, md.Data)
	}
}

//Block 区块内容测试
func (mysuit *MySuite) TestTestSdk_BlockData(c *C) {
	ut.Init(orgID)
	contractOwner := ut.DeployContract(c, contractName, orgID, contractMethods, contractMethods)
	test := NewTestObject(contractOwner)

	//Checking block height
	ut.AssertEquals(test.obj.sdk.Block().Height(), int64(1))
	setMystruct(test, contractOwner) // generate 8 blocks for each calling
	ut.AssertEquals(test.obj.sdk.Block().Height(), int64(9))
	ut.AssertEquals(test.obj.sdk.Block().ChainID(), "local")
	//其他接口没有测试的必要
}

func (mysuit *MySuite) TestTestSdk_IHelper(c *C) {
	ut.Init(orgID)
	contractOwner := ut.DeployContract(c, contractName, orgID, contractMethods, contractMethods)
	test := NewTestObject(contractOwner)

	setMystruct(test, contractOwner) // generate 8 blocks for each calling
	ut.AssertEquals(int64(9), test.obj.sdk.Block().Height())

	block := test.obj.sdk.Helper().BlockChainHelper().GetBlock(6)
	ut.AssertEquals(block.Height(), int64(6))
	ut.AssertEquals(block.NumTxs(), int32(1))
	//
}

//IAccount
func (mysuit *MySuite) TestTestSdk_IAccount(c *C) {
	ut.Init(orgID)
	contractOwner := ut.DeployContract(c, contractName, orgID, contractMethods, contractMethods)
	test := NewTestObject(contractOwner)

	//ut.AssertEquals(test.obj.sdk.Contract().Account(), "") //Address
	genowner := test.obj.sdk.Helper().GenesisHelper().Token().Owner()
	acc := test.obj.sdk.Message().Contract().Account()

	value := N(1E12)
	genowner.TransferByName("loc", acc.Address(), value)
	ut.Commit()
	//ut.AssertEquals(acc.Balance().String(), value.String())
	ut.AssertEquals(acc.BalanceOfName("loc").String(), value.String())
	ut.AssertEquals(acc.BalanceOfSymbol("loc").String(), value.String())
	ut.AssertEquals(acc.BalanceOfToken(test.obj.sdk.Helper().GenesisHelper().Token().Address()).String(), value.String())

	tacc := test.obj.sdk.Helper().AccountHelper().AccountOfPubKey(N(rand.Int63()).Value().Bytes())
	//transfer
	v1 := N(1E5)
	acc.Transfer(tacc.Address(), v1)
	ut.Commit()
	ut.AssertEquals(tacc.Balance().String(), v1.String())
	//transfer by token
	acc.TransferByToken(test.obj.sdk.Helper().GenesisHelper().Token().Address(), tacc.Address(), v1)
	ut.Commit()
	ut.AssertEquals(tacc.BalanceOfToken(test.obj.sdk.Helper().GenesisHelper().Token().Address()).String(), v1.MulI(2).String())

	//transfer by name
	acc.TransferByName(test.obj.sdk.Helper().GenesisHelper().Token().Name(), tacc.Address(), v1)
	ut.Commit()
	ut.AssertEquals(tacc.BalanceOfToken(test.obj.sdk.Helper().GenesisHelper().Token().Address()).String(), v1.MulI(3).String())

	//transfer by symbol
	acc.TransferBySymbol(test.obj.sdk.Helper().GenesisHelper().Token().Symbol(), tacc.Address(), v1)
	ut.Commit()
	ut.AssertEquals(tacc.BalanceOfToken(test.obj.sdk.Helper().GenesisHelper().Token().Address()).String(), v1.MulI(4).String())

}

//IAccountHelper
func (mysuit *MySuite) TestTestSdk_IAccountHelper(c *C) {
	ut.Init(orgID)
	contractOwner := ut.DeployContract(c, contractName, orgID, contractMethods, contractMethods)
	test := NewTestObject(contractOwner)

	//测试用例： 取创世owner,生成创世owner的账户，对比信息
	genowner := test.obj.sdk.Helper().GenesisHelper().Token().Owner()
	ut.Assert(genowner != nil)
	ut.Assert(genowner.Balance().IsLE(test.obj.sdk.Helper().GenesisHelper().Token().TotalSupply()))
	ut.AssertEquals(genowner.Address(), genowner)

	//	AccountOfPubKey(_pubkey PubKey) IAccount     //根据账户公钥构造账户信息对象
	// 测试用例： 根据上一步得到的创世账户的公钥生成一个新账户，对比账户信息，必须完全匹配
	acc := test.obj.sdk.Helper().AccountHelper().AccountOfPubKey(genowner.PubKey())
	ut.Assert(acc != nil)
	ut.AssertEquals(acc.Address(), genowner.Address())
	ut.AssertEquals(acc.PubKey(), genowner.PubKey())
	ut.AssertEquals(acc.Balance().String(), genowner.Balance().String())
}

//IBlockChainHelper
func (mysuit *MySuite) TestTestSdk_IBlockChainHelper(c *C) {
	ut.Init(orgID)
	contractOwner := ut.DeployContract(c, contractName, orgID, contractMethods, nil)
	test := NewTestObject(contractOwner)
	helper := test.obj.sdk.Helper().BlockChainHelper()

	//	CalcAccountFromPubKey(_pubkey []byte) Address                              //根据用户公钥计算账户地址
	addr1 := helper.CalcAccountFromPubKey([]byte("tt123456789012345678901234567890"))
	//生成地址时截取32个字符长度
	addr2 := helper.CalcAccountFromPubKey([]byte("tt1234567890123456789012345678901234567890"))
	ut.AssertEquals(addr1, addr2)

	//	CalcAccountFromName(_name string) Address                                  //根据合约名称计算合约的账户地址
	//  检查本合约的账户地址
	accaddr := helper.CalcAccountFromName(contractName, "")
	ut.AssertEquals(accaddr, test.obj.sdk.Message().Contract().Account().Address())
	//  输入Name 为空，依然可以生成一个地址
	ut.Assert(helper.CalcAccountFromName("", "") != "")

	//	CalcContractAddress(_name string, _version string, _owner Address) Address //根据合约名称、版本与所有者地址计算合约地址
	//  检查本合约的地址
	ver := test.obj.sdk.Message().Contract().Version()
	owner := test.obj.sdk.Message().Contract().Owner()
	ctraddr := helper.CalcContractAddress(contractName, ver, owner.Address())
	ut.AssertEquals(test.obj.sdk.Message().Contract().Address(), ctraddr)
	//  参数为空，依然可以生成一个地址
	ut.Assert(helper.CalcContractAddress("", "", "") != "")

	//	CheckAddress(_addr Address) Error                                          //根据chainID检查地址是否合法

	//	GetBlock(_height int64) IBlock                                             //根据高度读取区块信息
	block := helper.GetBlock(1)
	ut.AssertEquals(block.Height(), int64(1))
	//当前块
	curHeight := test.obj.sdk.Block().Height()
	curBlock := helper.GetBlock(curHeight)
	ut.AssertEquals(curBlock.Height(), curHeight)
	//下一个块（不存在）
	ut.Assert(helper.GetBlock(curHeight+1) == nil)
}

//IContractHelper
func (mysuit *MySuite) TestTestSdk_IContractHelper(c *C) {
	ut.Init(orgID)
	contractOwner := ut.DeployContract(c, contractName, orgID, contractMethods, contractMethods)
	test := NewTestObject(contractOwner)

	helper := test.obj.sdk.Helper().ContractHelper()

	//	ContractOfAddress(_addr Address) IContract    //根据合约地址构造合约信息对象
	ic1 := helper.ContractOfAddress(test.obj.sdk.Message().Contract().Address())

	//	ContractOfName(_name string) IContract        //根据合约名字返回当前有效合约对象
	ic2 := helper.ContractOfName(contractName)
	ut.AssertEquals(ic1.Address(), ic2.Address())
	ut.AssertEquals(ic1.Owner().Address(), ic2.Owner().Address())
	ut.AssertEquals(ic1.CodeHash().String(), ic2.CodeHash().String())

	//不存在的合约名称
	ut.Assert(helper.ContractOfName("") == nil)
	ut.Assert(helper.ContractOfAddress("") == nil)
	ut.Assert(helper.ContractOfToken("") == nil)
	ut.Assert(helper.ContractOfName("123") == nil)
	ut.Assert(helper.ContractOfName("adbcdeddddddddddddd") == nil)

	//	ContractOfToken(_tokenAddr Address) IContract //根据代币地址构造合约信息对象（当前区块可用）

}

//IReceiptHelper
func (mysuit *MySuite) TestTestSdk_IReceiptHelper(c *C) {
	ut.Init(orgID)
	contractOwner := ut.DeployContract(c, contractName, orgID, contractMethods, contractMethods)
	test := NewTestObject(contractOwner)

	helper := test.obj.sdk.Helper().ReceiptHelper()

	//	Emit(_interface interface{}) //发送一个事件
	type myreceipt struct {
		owner string
	}
	r := myreceipt{owner: "ddd"}
	helper.Emit(r)
	//todo: 补充测试用例
	//	ParseToTransfer(_interface interface{}) *receipt.Transfer
}

//IToken
func (mysuit *MySuite) TestTestSdk_IToken(c *C) {
	ut.Init(orgID)
	contractOwner := ut.DeployContract(c, contractName, orgID, contractMethods, nil)
	test := NewTestObject(contractOwner)

	token := test.obj.sdk.Helper().TokenHelper().TokenOfContract(test.obj.sdk.Message().Contract().Address())

	//  合约没有注册代币，token 为空
	ut.Assert(token == nil)

	// todo: 注册代币，测试代币数据
	tokenName := "testcoin"
	tokenSymbol := "tcn"
	totalSupply := N(1E15)

	icontract := test.obj.sdk.Message().Contract()
	newobject := test.run().setSender(test.obj.sdk.Helper().AccountHelper().AccountOf(icontract.Owner().Address()))
	newtoken := newobject.obj.sdk.Helper().TokenHelper().RegisterToken(tokenName, tokenSymbol, totalSupply, false, false)
	ut.Assert(newtoken != nil)

	//	Address() Address                         //代币地址
	ut.AssertEquals(newtoken.Address(), icontract.Address())
	//	Owner() Address                           //代币拥有者的账户地址
	ut.AssertEquals(newtoken.Owner().Address(), icontract.Owner().Address())
	//	Name() string                             //代币的名称
	//	Symbol() string                           //代币的符号
	//	TotalSupply() Number                      //代币的总供应量
	//	AddSupplyEnabled() bool                   //代币是否支持增发
	//	BurnEnabled() bool                        //代币是否支持燃烧
	//	GasPrice() int64                          //代币燃料价格
	//	SetOwner(_owner Address) Error            //设置代币拥有者的账户地址
	//	SetTotalSupply(_totalSupply Number) Error //设置代币的总供应量
	//	SetGasPrice(_gasPrice int64) Error        //设置代币燃料价格
}

//IGenesisHelper
func (mysuit *MySuite) TestTestSdk_IGenesisHelper(c *C) {
	ut.Init(orgID)
	contractOwner := ut.DeployContract(c, contractName, orgID, contractMethods, nil)
	test := NewTestObject(contractOwner)

	helper := test.obj.sdk.Helper().GenesisHelper()

	//	ChainId() string        //读取创世时的链ID
	ut.AssertEquals(helper.ChainID(), "local")
	//TODO
	//	Contracts() []IContract //读取创世合约信息

	//	Token() IToken          //读取创世通证（基础通证）的信息
}

//IStateHelper
func (mysuit *MySuite) TestTestSdk_IStateHelper(c *C) {
	ut.Init(orgID)
	contractOwner := ut.DeployContract(c, contractName, orgID, contractMethods, nil)
	test := NewTestObject(contractOwner)

	helper := test.obj.sdk.Helper().StateHelper()

	ut.Assert(helper.Check("/test") == false)
	//todo

	//  Check(_key string) bool // 判断指定的key对应的数据是否存在
	//	Get(_key string, _interface interface{}) interface{}   //从状态数据库中读取指定KEY对应的数据，不存在返回空
	//	GetEx(_key string, _interface interface{}) interface{} //从状态数据库中读取指定KEY对应的数据，不存在返回默认值
	//Set(_key string, _data interface{}) //向状态数据库设置指定KEY对应的数据
	//	//Memory cache McGet
	//	McGet(_key string, _interface interface{}) interface{}   //从状态数据库或内存缓存中读取指定KEY对应的数据，不存在返回空
	//	McGetEx(_key string, _interface interface{}) interface{} //从状态数据库或内存缓存中读取指定KEY对应的数据，不存在返回默认值
	//	McSet(_key string, _data interface{}) //向状态数据库和内存缓存设置指定KEY对应的数据
	//	McClear(_key string)
}

//ITokenHelper
func (mysuit *MySuite) TestTestSdk_ITokenHelper(c *C) {
	ut.Init(orgID)
	contractOwner := ut.DeployContract(c, contractName, orgID, contractMethods, nil)
	test := NewTestObject(contractOwner)

	helper := test.obj.sdk.Helper().TokenHelper()
	//  代码名称为空，取本币
	ut.Assert(helper.TokenOfName("") != nil)
	curToken := "testcoin"
	curSymbol := "tcn"
	// 注册代币，测试代币数据
	var testcases = []struct {
		tokenName        string
		tokenSymbol      string
		totalSupply      Number
		addSupplyEnabled bool
		burnEnabled      bool
		expectedError    uint32
	}{
		//名称为空
		{"", curToken, N(1E14), true, true, types.ErrInvalidParameter},
		//名称为一个字符
		{"a", curToken, N(1E14), true, true, types.ErrInvalidParameter},
		//名称为两个字符
		{"tc", curToken, N(1E14), true, true, types.ErrInvalidParameter},
		//名称大于40字符
		{"testcointestcointestcointestcointestcointestcoin", "testcoin", N(1E14), true, true, types.ErrInvalidParameter},
		//符号为空
		{curToken, "", N(1E14), true, true, types.ErrInvalidParameter},
		//符号为一个字符
		{curToken, "t", N(1E14), true, true, types.ErrInvalidParameter},
		//符号为两个字符
		{curToken, "tc", N(1E14), true, true, types.ErrInvalidParameter},
		//符号多余20个字符
		{curToken, "testcointestcointestcoin", N(1E14), true, true, types.ErrInvalidParameter},
		//供应量为1
		{curToken, curSymbol, N(1), true, true, types.ErrInvalidParameter},
		//供应量为1E8 < 1E9
		{curToken, curSymbol, N(1E8), true, true, types.ErrInvalidParameter},
		//供应量大于最大限制( 1<<256 cong)
		{curToken, curSymbol, N(1).Lsh(260), true, true, types.ErrInvalidParameter},
		//供应量为0,抢注代币
		{curToken, curSymbol, N(0), true, true, types.CodeOK},
		//正常用例
		{curToken, curSymbol, N(1E15), true, true, types.CodeOK},

		//为合约注册多个代码，失败
		{curToken + curToken, curSymbol + curSymbol, N(1E15), true, true, types.ErrNoAuthorization},
	}
	icontract := test.obj.sdk.Message().Contract()
	newobject := test.run().setSender(test.obj.sdk.Helper().AccountHelper().AccountOf(icontract.Owner().Address()))

	//	RegisterToken(_name, _symbol string, totalSupply Number, addSupplyEnabled, burnEnabled bool) (IToken, Error) //注册一个新的代币
	for _, c := range testcases {
		newtoken := newobject.obj.sdk.Helper().TokenHelper().RegisterToken(c.tokenName, c.tokenSymbol, c.totalSupply, c.addSupplyEnabled, c.burnEnabled)
		if c.expectedError == types.CodeOK {
			ut.Assert(newtoken != nil)
		} else {
			ut.Assert(newtoken == nil)
		}
	}
	gen := newobject.obj.sdk.Helper().GenesisHelper()
	//测试用例：创世币，本合约代币
	var tokencases = []struct {
		tokenName    string
		symbol       string
		addr         types.Address
		contractName string
		itoken       sdk.IToken
	}{
		//创世币
		{gen.Token().Name(), gen.Token().Symbol(), gen.Token().Address(), "token-basic", gen.Token()},
		//本合约代币
		{curToken, curSymbol, newobject.obj.sdk.Helper().TokenHelper().Token().Address(), contractName, newobject.obj.sdk.Helper().TokenHelper().Token()},
	}

	for _, c := range tokencases {
		//	TokenOfAddress(_tokenAddr Address) IToken  //根据代币地址获取代币或基础通证的信息
		testToken := helper.TokenOfAddress(c.addr)
		checkTokens(testToken, c.itoken)

		//	TokenOfName(_name string) IToken  //根据代币名称获取代币或基础通证的信息
		testToken = helper.TokenOfName(c.tokenName)
		checkTokens(testToken, c.itoken)

		//	TokenOfSymbol(_symbol string) IToken    //根据代币符号获取代币或基础通证的信息
		testToken = helper.TokenOfSymbol(gen.Token().Symbol())
		checkTokens(testToken, c.itoken)

		//	TokenOfContract(_contractAddr Address) IToken   //根据合约地址获取代币或基础通证的信息
		tbc := newobject.obj.sdk.Helper().ContractHelper().ContractOfName(c.contractName)
		testToken = helper.TokenOfContract(tbc.Address())
		checkTokens(testToken, c.itoken)
	}

	// 不存在的代币
	ut.Assert(helper.TokenOfAddress(gen.Token().Owner().Address()) == nil)
	ut.Assert(helper.TokenOfName(curToken+curSymbol) == nil)
	ut.Assert(helper.TokenOfSymbol(curToken+curSymbol) == nil)
	ut.Assert(helper.TokenOfContract(newobject.obj.sdk.Helper().ContractHelper().ContractOfName("myplayerbook").Address()) == nil)
}

func checkTokens(a, b sdk.IToken) {
	ut.Assert(a != nil)
	ut.Assert(b != nil)
	ut.AssertEquals(a.Address(), b.Address())
	ut.AssertEquals(a.Owner(), b.Owner())
	ut.AssertEquals(a.TotalSupply(), b.TotalSupply())
	ut.AssertEquals(a.Name(), b.Name())
	ut.AssertEquals(a.GasPrice(), b.GasPrice())
	ut.AssertEquals(a.BurnEnabled(), b.BurnEnabled())
	ut.AssertEquals(a.AddSupplyEnabled(), b.AddSupplyEnabled())
}

func (mysuit *MySuite) TestTestSdk_SetDatai16(c *C) {
	ut.Init(orgID)
	contractOwner := ut.DeployContract(c, contractName, orgID, contractMethods, nil)
	test := NewTestObject(contractOwner)

	test.SetDatai16(100)
}
func (mysuit *MySuite) TestTestSdk_GetDatai16(c *C) {
	ut.Init(orgID)
	contractOwner := ut.DeployContract(c, contractName, orgID, contractMethods, nil)
	test := NewTestObject(contractOwner)

	test.GetDatai16()
}
func (mysuit *MySuite) TestTestSdk_SetOwner(c *C) {
	ut.Init(orgID)
	contractOwner := ut.DeployContract(c, contractName, orgID, contractMethods, nil)
	test := NewTestObject(contractOwner)

	test.SetOwner("")
}
func (mysuit *MySuite) TestTestSdk_SetMydata(c *C) {
	ut.Init(orgID)
	contractOwner := ut.DeployContract(c, contractName, orgID, contractMethods, nil)
	test := NewTestObject(contractOwner)

	v := Mystruct{}
	test.SetMydata(v)
}
