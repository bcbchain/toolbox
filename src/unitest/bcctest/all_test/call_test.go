package all

import (
	"encoding/json"
	"fmt"
	"testing"
	"unitest/bcctest/common"
)

func TestDemo1(t *testing.T) {

	var tests = []struct {
		Method string
		param  string
		Msg    string
		Desc   string
	}{
		{"EchoAddress", "localL9BzYNYns5VCRaJgfHEBJLzS1bhpHjx7j", "", "Success"},
		{"EchoHash", "1b4e92bc742dc29076087469d49d3384216d5f79532a9fb4dc7b1933bd9a4262", "", "Success"},
		{"EchoHexBytes", "AAAAAAAAAA", "", "Success"},
		{"EchoPubKey", "01bd6c29d63f5f32aa33955f26a28459988edea4de517f77372e77db33958e6e", "", "Success"},
		{"EchoNumber", "66666666666666", "", "Success"},
		{"EchoInt", "66666666", "", "Success"},
		{"EchoInt8", "66", "", "Success"},
		{"EchoInt16", "66", "", "Success"},
		{"EchoInt32", "66", "", "Success"},
		{"EchoInt64", "66", "", "Success"},
		{"EchoUint", "66", "", "Success"},
		{"EchoUint8", "66", "", "Success"},
		{"EchoUint16", "66", "", "Success"},
		{"EchoUint32", "66", "", "Success"},
		{"EchoUint64", "66", "", "Success"},
		{"EchoBool", "true", "", "Success"},

		{"EchoByte", "\"0x97\"", "", "Success"},
		{"EchoBytes", "0x97", "", "Success"},
		{"EchoMap", "{\"test\":123}", "", "Success"},
		{"EchoMap1", "{\"123\":{\"test\":\"localL9BzYNYns5VCRaJgfHEBJLzS1bhpHjx7j\"}}", "", "Success"},

		{"EchoMap2", "{\"123\":{\"1\":\"1b4e92bc742dc29076087469d49d3384216d5f79532a9fb4dc7b1933bd9a4262\"}}", "", "Success"},
		{"EchoMap3", "{\"123\":{\"1\":\"AAAAAAAAAA\"}}", "", "Success"},

		{"EchoMap4", "{\"true\":{\"0x11\":\"1b4e92bc742dc29076087469d49d3384216d5f79532a9fb4dc7b1933bd9a4262\"}}", "", "Success"},
		{"EchoMap5", "{\"true\":{\"666\":true}}", "", "Success"},

		{"EchoMap6", "{\"0x97\":{\"test\":\"666\"}}", "", "Success"},
		{"EchoMap7", "{\"test\":{\"localL9BzYNYns5VCRaJgfHEBJLzS1bhpHjx7j\":\"0x97\"}}", "", "Success"},
		{"EchoMap8", "{\"localL9BzYNYns5VCRaJgfHEBJLzS1bhpHjx7j\":{\"66\":\"test\"}}", "", "Success"},
	}
	success := common.Success
	common.PrintTitle("TestDemo1")

	for index, test := range tests {
		result, err := common.Call(name, password, orgName, contractNamedemo1, test.Method, "", test.param, "|", "", gasLimit, "", chainID, keystorePath)
		common.PrintResult(err)
		jsIndent := []byte{}
		if result != nil {
			jsIndent, _ = json.MarshalIndent(&result.Data, "", "\t")
		}
		common.PrintCase(index, test.Desc, test.Msg, "Data:"+string(jsIndent), err)

		fmt.Println(test.Desc)
		common.AssertSuccess(common.Success-success, len(tests))
	}

}
