package blockchain

import (
	"github.com/bcbchain/toolbox/relay/common"
)

func QueryOpenURLs(chainID string) ([]string, error) {

	config := common.GetConfig()
	var openUrls []string

	err := common.DoHttpQueryAndParse(config.MainChainUrl, "/sidechain/"+chainID+"/openurls", &openUrls)
	if err != nil {
		return nil, err
	}

	return openUrls, nil
}
