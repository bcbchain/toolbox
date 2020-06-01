package db

import (
	"github.com/bcbchain/bclib/bcdb"
	"github.com/bcbchain/bclib/jsoniter"
	"github.com/bcbchain/tendermint/types"
	"sync"
)

var (
	sdb  *bcdb.GILevelDB
	once sync.Once
)

func InitDB(dbName, dbPort string) (err error) {
	once.Do(func() {
		sdb, err = bcdb.OpenDB(dbName, "", dbPort)
	})

	return err
}

func Close() {
	sdb.Close()
}

func GetFirstHeight() int64 {
	var h int64
	get(keyFirstHeight(), &h)
	return h
}

func SetFirstHeight(h int64) {
	setSync(keyFirstHeight(), h)
}

func GetLastHeight() int64 {
	var h int64
	get(keyLastHeight(), &h)
	return h
}

func SetLastHeight(h int64) {
	setSync(keyLastHeight(), h)
}

func SetHeader(h int64, header *types.Header) {
	setSync(keyOfHeader(h), header)
}

func GetHeader(h int64) *types.Header {
	header := new(types.Header)
	get(keyOfHeader(h), header)
	return header
}

func HasHeader(h int64) bool {
	header := new(types.Header)
	get(keyOfHeader(h), header)
	if header.ChainID != "" {
		return true
	} else {
		return false
	}
}

func SetTx(h, index int64, tx string) {
	setSync(keyOfTx(h, index), tx)
}

func HasTxWithHeight(h int64) bool {
	tx := ""
	get(keyOfTx(h, 0), &tx)
	return tx != ""
}

func GetTx(h int64) []string {

	result := make([]string, 0)
	for i := int64(0); i < 50000; i++ {
		tx := ""
		get(keyOfTx(h, i), &tx)
		if tx != "" {
			result = append(result, tx)
		} else {
			break
		}
	}

	return result
}

func setSync(key string, value interface{}) {
	v, err := jsoniter.Marshal(value)
	if err != nil {
		panic(err)
	}
	if err = sdb.SetSync([]byte(key), v); err != nil {
		panic(err)
	}
}

// objPoint must be point
func get(key string, objPoint interface{}) {
	value, err := sdb.Get([]byte(key))
	if err != nil {
		panic(err)
	}

	if len(value) == 0 {
		return
	}

	err = jsoniter.Unmarshal(value, objPoint)
	if err != nil {
		panic(err)
	}
}
