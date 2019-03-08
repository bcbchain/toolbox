package pvar

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestCreate(t *testing.T) {
	typeStrList := "bn.Number,int,uint,bool,byte,[]byte,float64,map[string]uint64,map[bool]map[bn.Number]string,[]string,string"
	valueStrList := `123333221,-123,1234,true,61,616263,123.456,{"test":123\,"test2":456},{"true":{"123":"test"}},["123"\,"456"],"str\,"\\ing"`

	splitTypes := strings.Split(typeStrList, ",")

	splitValues, err := SplitValues(valueStrList)
	if err != nil {
		t.Error(err)
		return
	}

	if len(splitValues) != len(splitTypes) {
		t.Error(fmt.Sprintf("values count=%d not equel types count=%d", len(splitValues), len(splitTypes)))
		return
	}

	for index, typeStr := range splitTypes {
		v, err := Create(typeStr, splitValues[index])
		if err != nil {
			t.Error(err)
			return
		}

		fmt.Println(typeStr, splitValues[index])
		fmt.Println(reflect.TypeOf(v).Kind(), v)
	}
}
