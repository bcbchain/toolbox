package scan

import (
	"fmt"
	"strconv"
	"strings"
)

type Result struct {
	ScannerID   ScannerID
	BeginHeight int64
	EndHeight   int64
	Err         error
}

type ScannerID string

func calcScannerID(url string, index int) ScannerID {
	return ScannerID(fmt.Sprintf("%s--%d", url, index))
}

func (id *ScannerID) string() string {
	return string(*id)
}

func (id *ScannerID) getURL() string {
	return strings.Split(id.string(), "--")[0]
}

func (id *ScannerID) getIndex() int {
	temp := strings.Split(id.string(), "--")
	index, _ := strconv.Atoi(temp[1])
	return index
}
