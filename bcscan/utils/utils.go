package utils

import (
	"fmt"
	"strings"
)

const (
	FilterTx     = "tx"
	FilterHeader = "header"
)

func ValidFilter(items string) ([]string, error) {
	if len(items) == 0 {
		return []string{FilterTx, FilterHeader}, nil
	}

	filterMap := make(map[string]struct{})
	itemList := strings.Split(items, ",")
	for _, item := range itemList {
		if item != FilterTx && item != FilterHeader {
			return nil, fmt.Errorf("invalid filter type: %s", item)
		}
		filterMap[item] = struct{}{}
	}

	result := make([]string, 0)
	for key := range filterMap {
		result = append(result, key)
	}
	return result, nil
}
