package cache

import (
	"cmd/bcc/common"
	"common/bcdb"
	"fmt"
	"github.com/tendermint/tendermint/rpc/core/types"
	"math"
	"time"
)

const (
	layout = "2006-01-02 15:04:05 +0000 UTC"
)

var (
	db     *bcdb.GILevelDB
	addrS  []string
	dbName = ".bccCache"
	dbIP   = "127.0.0.1"
	dbPort = "55678"
)

// KeyOfHeight return key of height for access time
func KeyOfHeight(height int64) []byte {
	return []byte(fmt.Sprintf("/%d", height))
}

// SetAddrList set nodes addresses
func SetAddrList(addrList []string) {
	addrS = addrList
}

// BinarySearchEx binary search algorithm with interval
func BinarySearchEx(start, end, interval int64, t time.Time) (height int64) {
	if start != 0 && start >= end {
		if CompareWithTime(start, t, false) != 0 {
			return
		}
		return start
	}

	// set start with current block height
	if start == 0 {
		result := new(core_types.ResultABCIInfo)
		err := common.DoHttpRequestAndParse(addrS, "abci_info", map[string]interface{}{}, result)
		if err != nil {
			panic(err)
		}
		start = result.Response.LastBlockHeight

		// check latest block time
		if r := CompareWithTime(start, t, false); r == -1 {
			return
		} else if r == 0 {
			return end
		}

		// check oldest block time
		if r := CompareWithTime(1, t, true); r == 1 {
			return
		} else if r == 0 {
			return 1
		}
	}

	// reset start and end
	if end == 0 {
		end = start
		start = (start / interval) * interval
	}

	// compare time between start's time and t
	r := CompareWithTime(start, t, false)
	if r == 1 {
		if start != 1 {
			end = start
			start = 1
		}

		m := ((end - 1) / interval / 2) * interval
		r = CompareWithTime(m, t, true)
		if r == 1 {
			return BinarySearchEx(start, m, interval, t)
		} else if r == -1 {
			return BinarySearchEx(m, end, interval, t)
		} else {
			return m
		}
	} else if r == -1 {
		if end-start <= interval {
			return binarySearch(start, end, t)
		} else {
			m := start + ((end-start)/interval/2)*interval
			if m == start {
				m = start + interval
			}
			r = CompareWithTime(m, t, true)
			if r == 1 {
				return BinarySearchEx(start, m, interval, t)
			} else if r == -1 {
				return BinarySearchEx(m, end, interval, t)
			} else {
				return m
			}
		}
	} else {
		return start
	}
}

func binarySearch(start, end int64, t time.Time) (height int64) {
	if start > end {
		return
	}

	if start == end {
		return start
	} else if start+1 == end {
		return nearlyBetweenTwoHeight(start, end, t)
	} else {
		m := (start + end) / 2
		r := CompareWithTime(m, t, false)
		if r == 1 {
			return binarySearch(start, m, t)
		} else if r == -1 {
			return binarySearch(m, end, t)
		} else {
			return m
		}
	}
}

// CompareWithTime compare time between block time and t,
// then return 1 if block time is bigger, else return -1,
// block time equal t return 0
func CompareWithTime(h int64, t time.Time, bSave bool) int {
	ht := timeOfHeightFromBlock(h)

	if bSave {
		setTimeOfHeight(h, ht.String())
	}

	if ht.Sub(t) > 0 {
		return 1
	} else if ht.Sub(t) < 0 {
		return -1
	}

	return 0
}

func timeOfHeight(height int64) string {
	v, err := db.Get(KeyOfHeight(height))
	if err != nil {
		return ""
	}

	return string(v)
}

func setTimeOfHeight(height int64, t string) {
	err := db.Set(KeyOfHeight(height), []byte(t))
	if err != nil {
		panic(err)
	}
}

func timeOfHeightFromBlock(height int64) time.Time {
	var err error
	var t time.Time
	tStr := timeOfHeight(height)
	if len(tStr) == 0 {
		result := new(core_types.ResultBlock)
		err = common.DoHttpRequestAndParse(addrS, "block", map[string]interface{}{"height": height}, result)
		if err != nil {
			panic(err)
		}
		t = result.BlockMeta.Header.Time
	} else {
		t, err = time.ParseInLocation(layout, tStr, time.UTC)
		if err != nil {
			panic(err)
		}
	}

	return t
}

func nearlyBetweenTwoHeight(h1, h2 int64, t time.Time) (h int64) {
	t1 := timeOfHeightFromBlock(h1)
	t2 := timeOfHeightFromBlock(h2)

	d1 := math.Abs(float64(t1.Sub(t).Nanoseconds()))
	d2 := math.Abs(float64(t2.Sub(t).Nanoseconds()))
	if d1 > d2 {
		return h2
	} else {
		return h1
	}
}
