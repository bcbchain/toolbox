package testblock

import (
	"blockchain/smcsdk/sdk/types"
	"blockchain/smcsdk/sdkimpl"
	"blockchain/smcsdk/sdkimpl/object"
)

func (t *TestBlock) testChainID() {

	cases := []struct {
		chainID string
		desc    string
		code    uint32
		errMsg  string
	}{
		{"", "正常用例-chainID为空", types.CodeOK, ""},
		{"ahshfiwehhfahsdfhasdlfaldsjhfoih", "正常用例-chainID为随机字符串", types.CodeOK, ""},
		{"bcbtest", "正常用例", types.CodeOK, ""},
		{"local", "正常用例", types.CodeOK, ""},
	}

	for i, c := range cases {
		printTestCase(i, c.desc)
		b := object.NewBlock(t.sdk, c.chainID,
			t.sdk.Block().Version(), t.sdk.Block().LastBlockHash(), t.sdk.Block().DataHash(),
			t.sdk.Block().Height(), t.sdk.Block().Time(), t.sdk.Block().NumTxs(),
			t.sdk.Block().ProposerAddress(), t.sdk.Block().RewardAddress(), t.sdk.Block().RandomNumber(),
			t.sdk.Block().LastBlockHash(), t.sdk.Block().LastCommitHash(), t.sdk.Block().LastAppHash(),
			t.sdk.Block().LastFee())

		smc := t.sdk.(*sdkimpl.SmartContract)
		smc.SetBlock(b)
		Assert(t.sdk.Block().ChainID() == c.chainID)
		printPass()
	}
}

func (t *TestBlock) testBlockHash() {
	cases := []struct {
		blockHash types.Hash
		desc      string
		code      uint32
		errMsg    string
	}{
		{types.Hash{}, "正常用例-blockHash为空", types.CodeOK, ""},
		{types.Hash([]byte(`{"chainID":"local","blockHash":"663666426932665034586B70685A5A323757424F786845473079493D","height":1,"time":1542436677,"numTxs":1,"dataHash":"4369496472305A546757635671336C3131326D31773539684557593D","proposerAddress":"localCUh7Zsb7PBgLwHJVok2QaMhbW64HNK4FU","rewardAddress":"localCUh7Zsb7PBgLwHJVok2QaMhbW64HNK4FU","randomNumber":"596D4E694E3074346546704E64314E3057556448626B74704D33684F546D3035625735744E33426B563352474F576445","version":"1.0","lastBlockHash":"4369496472305A546757635671336C3131326D31773539684557593D","lastCommitHash":"4D31724C4573745A68314B30624A644D6B4A4B53625774324450673D","lastAppHash":"36324446453643413937353539313437324435323143413943303845413243453242444646333546363231363230393132393937324636414133334633314234","lastFee":1500000}`)), "正常用例", types.CodeOK, ""},
	}

	for i, c := range cases {
		printTestCase(i, c.desc)
		b := object.NewBlock(t.sdk, t.sdk.Block().ChainID(),
			t.sdk.Block().Version(), c.blockHash, t.sdk.Block().DataHash(),
			t.sdk.Block().Height(), t.sdk.Block().Time(), t.sdk.Block().NumTxs(),
			t.sdk.Block().ProposerAddress(), t.sdk.Block().RewardAddress(), t.sdk.Block().RandomNumber(),
			t.sdk.Block().LastBlockHash(), t.sdk.Block().LastCommitHash(), t.sdk.Block().LastAppHash(),
			t.sdk.Block().LastFee())

		smc := t.sdk.(*sdkimpl.SmartContract)
		smc.SetBlock(b)
		for index, v := range t.sdk.Block().BlockHash() {
			Assert(v == c.blockHash[index])
		}

		printPass()
	}
}

func (t *TestBlock) testHeight() {
	cases := []struct {
		height int64
		desc   string
		code   uint32
		errMsg string
	}{
		{0, "正常用例", types.CodeOK, ""},
		{1, "正常用例", types.CodeOK, ""},
		{-2, "正常用例", types.CodeOK, ""},
		{200, "正常用例", types.CodeOK, ""},
	}

	for i, c := range cases {
		printTestCase(i, c.desc)
		b := object.NewBlock(t.sdk, t.sdk.Block().ChainID(),
			t.sdk.Block().Version(), t.sdk.Block().BlockHash(), t.sdk.Block().DataHash(),
			c.height, t.sdk.Block().Time(), t.sdk.Block().NumTxs(),
			t.sdk.Block().ProposerAddress(), t.sdk.Block().RewardAddress(), t.sdk.Block().RandomNumber(),
			t.sdk.Block().LastBlockHash(), t.sdk.Block().LastCommitHash(), t.sdk.Block().LastAppHash(),
			t.sdk.Block().LastFee())

		smc := t.sdk.(*sdkimpl.SmartContract)
		smc.SetBlock(b)
		Assert(t.sdk.Block().Height() == c.height)

		printPass()
	}
}

func (t *TestBlock) testTime() {
	cases := []struct {
		time   int64
		desc   string
		code   uint32
		errMsg string
	}{
		{0, "正常用例", types.CodeOK, ""},
		{1, "正常用例", types.CodeOK, ""},
		{100000, "正常用例", types.CodeOK, ""},
		{10000000000000, "正常用例", types.CodeOK, ""},
		{-100, "正常用例", types.CodeOK, ""},
	}

	for i, c := range cases {
		printTestCase(i, c.desc)
		b := object.NewBlock(t.sdk, t.sdk.Block().ChainID(),
			t.sdk.Block().Version(), t.sdk.Block().BlockHash(), t.sdk.Block().DataHash(),
			t.sdk.Block().Height(), c.time, t.sdk.Block().NumTxs(),
			t.sdk.Block().ProposerAddress(), t.sdk.Block().RewardAddress(), t.sdk.Block().RandomNumber(),
			t.sdk.Block().LastBlockHash(), t.sdk.Block().LastCommitHash(), t.sdk.Block().LastAppHash(),
			t.sdk.Block().LastFee())

		smc := t.sdk.(*sdkimpl.SmartContract)
		smc.SetBlock(b)
		Assert(t.sdk.Block().Time() == c.time)

		printPass()
	}
}

func (t *TestBlock) testNow() {
	cases := []struct {
		now    int64
		desc   string
		code   uint32
		errMsg string
	}{
		{0, "正常用例", types.CodeOK, ""},
		{1, "正常用例", types.CodeOK, ""},
		{100000, "正常用例", types.CodeOK, ""},
		{10000000000000, "正常用例", types.CodeOK, ""},
		{-100, "正常用例", types.CodeOK, ""},
	}

	for i, c := range cases {
		printTestCase(i, c.desc)
		b := object.NewBlock(t.sdk, t.sdk.Block().ChainID(),
			t.sdk.Block().Version(), t.sdk.Block().BlockHash(), t.sdk.Block().DataHash(),
			t.sdk.Block().Height(), c.now, t.sdk.Block().NumTxs(),
			t.sdk.Block().ProposerAddress(), t.sdk.Block().RewardAddress(), t.sdk.Block().RandomNumber(),
			t.sdk.Block().LastBlockHash(), t.sdk.Block().LastCommitHash(), t.sdk.Block().LastAppHash(),
			t.sdk.Block().LastFee())

		smc := t.sdk.(*sdkimpl.SmartContract)
		smc.SetBlock(b)
		Assert(t.sdk.Block().Time() == c.now)

		printPass()
	}
}

func (t *TestBlock) testNumTxs() {
	cases := []struct {
		numTxs int32
		desc   string
		code   uint32
		errMsg string
	}{
		{0, "正常用例", types.CodeOK, ""},
		{1, "正常用例", types.CodeOK, ""},
		{100000, "正常用例", types.CodeOK, ""},
		{9000000, "正常用例", types.CodeOK, ""},
		{-100, "正常用例", types.CodeOK, ""},
	}

	for i, c := range cases {
		printTestCase(i, c.desc)
		b := object.NewBlock(t.sdk, t.sdk.Block().ChainID(),
			t.sdk.Block().Version(), t.sdk.Block().BlockHash(), t.sdk.Block().DataHash(),
			t.sdk.Block().Height(), t.sdk.Block().Time(), c.numTxs,
			t.sdk.Block().ProposerAddress(), t.sdk.Block().RewardAddress(), t.sdk.Block().RandomNumber(),
			t.sdk.Block().LastBlockHash(), t.sdk.Block().LastCommitHash(), t.sdk.Block().LastAppHash(),
			t.sdk.Block().LastFee())

		smc := t.sdk.(*sdkimpl.SmartContract)
		smc.SetBlock(b)
		Assert(t.sdk.Block().NumTxs() == c.numTxs)

		printPass()
	}
}

func (t *TestBlock) testDataHash() {
	cases := []struct {
		dataHash types.Hash
		desc     string
		code     uint32
		errMsg   string
	}{
		{types.Hash{}, "正常用例", types.CodeOK, ""},
	}

	for i, c := range cases {
		printTestCase(i, c.desc)
		b := object.NewBlock(t.sdk, t.sdk.Block().ChainID(),
			t.sdk.Block().Version(), t.sdk.Block().BlockHash(), c.dataHash,
			t.sdk.Block().Height(), t.sdk.Block().Time(), t.sdk.Block().NumTxs(),
			t.sdk.Block().ProposerAddress(), t.sdk.Block().RewardAddress(), t.sdk.Block().RandomNumber(),
			t.sdk.Block().LastBlockHash(), t.sdk.Block().LastCommitHash(), t.sdk.Block().LastAppHash(),
			t.sdk.Block().LastFee())

		smc := t.sdk.(*sdkimpl.SmartContract)
		smc.SetBlock(b)
		for index, v := range t.sdk.Block().DataHash() {
			Assert(v == c.dataHash[index])
		}

		printPass()
	}
}

func (t *TestBlock) testProposerAddress() {
	cases := []struct {
		address types.Address
		desc    string
		code    uint32
		errMsg  string
	}{
		{"", "正常用例", types.CodeOK, ""},
	}

	for i, c := range cases {
		printTestCase(i, c.desc)
		b := object.NewBlock(t.sdk, t.sdk.Block().ChainID(),
			t.sdk.Block().Version(), t.sdk.Block().BlockHash(), t.sdk.Block().DataHash(),
			t.sdk.Block().Height(), t.sdk.Block().Time(), t.sdk.Block().NumTxs(),
			c.address, t.sdk.Block().RewardAddress(), t.sdk.Block().RandomNumber(),
			t.sdk.Block().LastBlockHash(), t.sdk.Block().LastCommitHash(), t.sdk.Block().LastAppHash(),
			t.sdk.Block().LastFee())

		smc := t.sdk.(*sdkimpl.SmartContract)
		smc.SetBlock(b)
		Assert(t.sdk.Block().ProposerAddress() == c.address)

		printPass()
	}
}

func (t *TestBlock) testRewardAddress() {
	cases := []struct {
		address types.Address
		desc    string
		code    uint32
		errMsg  string
	}{
		{"", "正常用例", types.CodeOK, ""},
	}

	for i, c := range cases {
		printTestCase(i, c.desc)
		b := object.NewBlock(t.sdk, t.sdk.Block().ChainID(),
			t.sdk.Block().Version(), t.sdk.Block().BlockHash(), t.sdk.Block().DataHash(),
			t.sdk.Block().Height(), t.sdk.Block().Time(), t.sdk.Block().NumTxs(),
			t.sdk.Block().ProposerAddress(), c.address, t.sdk.Block().RandomNumber(),
			t.sdk.Block().LastBlockHash(), t.sdk.Block().LastCommitHash(), t.sdk.Block().LastAppHash(),
			t.sdk.Block().LastFee())

		smc := t.sdk.(*sdkimpl.SmartContract)
		smc.SetBlock(b)
		Assert(t.sdk.Block().RewardAddress() == c.address)

		printPass()
	}
}

func (t *TestBlock) testRandomNumber() {
	cases := []struct {
		randNumber types.HexBytes
		desc       string
		code       uint32
		errMsg     string
	}{
		{types.HexBytes{}, "正常用例", types.CodeOK, ""},
	}

	for i, c := range cases {
		printTestCase(i, c.desc)
		b := object.NewBlock(t.sdk, t.sdk.Block().ChainID(),
			t.sdk.Block().Version(), t.sdk.Block().BlockHash(), t.sdk.Block().DataHash(),
			t.sdk.Block().Height(), t.sdk.Block().Time(), t.sdk.Block().NumTxs(),
			t.sdk.Block().ProposerAddress(), t.sdk.Block().RewardAddress(), c.randNumber,
			t.sdk.Block().LastBlockHash(), t.sdk.Block().LastCommitHash(), t.sdk.Block().LastAppHash(),
			t.sdk.Block().LastFee())

		smc := t.sdk.(*sdkimpl.SmartContract)
		smc.SetBlock(b)
		for index, v := range t.sdk.Block().RandomNumber() {
			Assert(v == c.randNumber[index])
		}

		printPass()
	}
}

func (t *TestBlock) testVersion() {
	cases := []struct {
		version string
		desc    string
		code    uint32
		errMsg  string
	}{
		{"", "正常用例", types.CodeOK, ""},
	}

	for i, c := range cases {
		printTestCase(i, c.desc)
		b := object.NewBlock(t.sdk, t.sdk.Block().ChainID(),
			c.version, t.sdk.Block().BlockHash(), t.sdk.Block().DataHash(),
			t.sdk.Block().Height(), t.sdk.Block().Time(), t.sdk.Block().NumTxs(),
			t.sdk.Block().ProposerAddress(), t.sdk.Block().RewardAddress(), t.sdk.Block().RandomNumber(),
			t.sdk.Block().LastBlockHash(), t.sdk.Block().LastCommitHash(), t.sdk.Block().LastAppHash(),
			t.sdk.Block().LastFee())

		smc := t.sdk.(*sdkimpl.SmartContract)
		smc.SetBlock(b)
		Assert(t.sdk.Block().Version() == c.version)

		printPass()
	}
}

func (t *TestBlock) testLastBlockHash() {
	cases := []struct {
		lastBlockHash types.Hash
		desc          string
		code          uint32
		errMsg        string
	}{
		{types.Hash{}, "正常用例", types.CodeOK, ""},
	}

	for i, c := range cases {
		printTestCase(i, c.desc)
		b := object.NewBlock(t.sdk, t.sdk.Block().ChainID(),
			t.sdk.Block().Version(), t.sdk.Block().BlockHash(), t.sdk.Block().DataHash(),
			t.sdk.Block().Height(), t.sdk.Block().Time(), t.sdk.Block().NumTxs(),
			t.sdk.Block().ProposerAddress(), t.sdk.Block().RewardAddress(), t.sdk.Block().RandomNumber(),
			c.lastBlockHash, t.sdk.Block().LastCommitHash(), t.sdk.Block().LastAppHash(),
			t.sdk.Block().LastFee())

		smc := t.sdk.(*sdkimpl.SmartContract)
		smc.SetBlock(b)
		for index, v := range t.sdk.Block().LastBlockHash() {
			Assert(v == c.lastBlockHash[index])
		}

		printPass()
	}
}

func (t *TestBlock) testLastCommitHash() {
	cases := []struct {
		lastCommitHash types.Hash
		desc           string
		code           uint32
		errMsg         string
	}{
		{types.Hash{}, "正常用例", types.CodeOK, ""},
	}

	for i, c := range cases {
		printTestCase(i, c.desc)
		b := object.NewBlock(t.sdk, t.sdk.Block().ChainID(),
			t.sdk.Block().Version(), t.sdk.Block().BlockHash(), t.sdk.Block().DataHash(),
			t.sdk.Block().Height(), t.sdk.Block().Time(), t.sdk.Block().NumTxs(),
			t.sdk.Block().ProposerAddress(), t.sdk.Block().RewardAddress(), t.sdk.Block().RandomNumber(),
			t.sdk.Block().LastBlockHash(), c.lastCommitHash, t.sdk.Block().LastAppHash(),
			t.sdk.Block().LastFee())

		smc := t.sdk.(*sdkimpl.SmartContract)
		smc.SetBlock(b)
		for index, v := range t.sdk.Block().LastCommitHash() {
			Assert(v == c.lastCommitHash[index])
		}

		printPass()
	}
}

func (t *TestBlock) testLastAppHash() {
	cases := []struct {
		lastAppHash types.Hash
		desc        string
		code        uint32
		errMsg      string
	}{
		{types.Hash{}, "正常用例", types.CodeOK, ""},
	}

	for i, c := range cases {
		printTestCase(i, c.desc)
		b := object.NewBlock(t.sdk, t.sdk.Block().ChainID(),
			t.sdk.Block().Version(), t.sdk.Block().BlockHash(), t.sdk.Block().DataHash(),
			t.sdk.Block().Height(), t.sdk.Block().Time(), t.sdk.Block().NumTxs(),
			t.sdk.Block().ProposerAddress(), t.sdk.Block().RewardAddress(), t.sdk.Block().RandomNumber(),
			t.sdk.Block().LastBlockHash(), t.sdk.Block().LastCommitHash(), c.lastAppHash,
			t.sdk.Block().LastFee())

		smc := t.sdk.(*sdkimpl.SmartContract)
		smc.SetBlock(b)
		for index, v := range t.sdk.Block().LastAppHash() {
			Assert(v == c.lastAppHash[index])
		}

		printPass()
	}
}
func (t *TestBlock) testLastFee() {
	cases := []struct {
		fee    int64
		desc   string
		code   uint32
		errMsg string
	}{
		{0, "正常用例", types.CodeOK, ""},
	}

	for i, c := range cases {
		printTestCase(i, c.desc)
		b := object.NewBlock(t.sdk, t.sdk.Block().ChainID(),
			t.sdk.Block().Version(), t.sdk.Block().BlockHash(), t.sdk.Block().DataHash(),
			t.sdk.Block().Height(), t.sdk.Block().Time(), t.sdk.Block().NumTxs(),
			t.sdk.Block().ProposerAddress(), t.sdk.Block().RewardAddress(), t.sdk.Block().RandomNumber(),
			t.sdk.Block().LastBlockHash(), t.sdk.Block().LastCommitHash(), t.sdk.Block().LastAppHash(),
			c.fee)

		smc := t.sdk.(*sdkimpl.SmartContract)
		smc.SetBlock(b)
		Assert(t.sdk.Block().LastFee() == c.fee)

		printPass()
	}
}
