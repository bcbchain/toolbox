package pvar

import (
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/std"
	"common/jsoniter"
	"encoding/hex"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var (
	mapTypes = []string{"int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "float32", "float64", "bn.Number", "bool", "byte", "types.Hash", "types.HexBytes", "types.PubKey", "[]byte"}
)

// Create create receive type and value as string, then return variable that defined type and set value
func Create(method std.Method, params, splitBy string) (results []interface{}, err error) {
	//splitValues, err := splitValues(params, splitBy)
	//if err != nil {
	//	return
	//}
	splitValues := strings.Split(params, splitBy)

	leftBracketIndex := strings.Index(method.ProtoType, "(")
	rightBracketIndex := strings.Index(method.ProtoType, ")")
	splitTypes := strings.Split(method.ProtoType[leftBracketIndex+1:rightBracketIndex], ",")

	if len(splitTypes) != len(splitValues) {
		return nil, errors.New(fmt.Sprintf("expect parameter's count=%d, obtain count=%d", len(splitTypes), len(splitValues)))
	}

	for index, typeStr := range splitTypes {
		var v interface{}
		v, err = create(typeStr, splitValues[index])
		if err != nil {
			return
		}

		results = append(results, v)
	}

	return
}

// Create create receive type and value as string, then return variable that defined type and set value
func create(typeStr, valueStr string) (v interface{}, err error) {
	var initV interface{}
	// step 1
	initV, err = initVariableWithValue(typeStr, valueStr)
	if err != nil {
		return
	}

	// step 2
	v, err = createVariable(typeStr, initV, false)

	return
}

func initVariableWithValue(typeStr, valueStr string) (v interface{}, err error) {

	temp, err := varFromString(typeStr)
	if err != nil {
		return
	}
	v = temp.Interface()

	if typeStr == "string" ||
		typeStr == "types.Address" ||
		typeStr == "types.Hash" ||
		typeStr == "[]byte" ||
		typeStr == "types.HexBytes" ||
		typeStr == "types.PubKey" {
		valueStr = strings.Replace(valueStr, `"`, `\"`, -1)
		valueStr = "\"" + valueStr + "\""
	}

	err = jsoniter.Unmarshal([]byte(valueStr), &v)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func createVariable(typeStr string, value interface{}, isMapKey bool) (v interface{}, err error) {
	v, err = createBaseVar(typeStr, value, isMapKey)
	if err != nil {
		return
	}

	if v == nil {
		if strings.HasPrefix(typeStr, "map") {
			v, err = createMapVar(typeStr, value)
		} else if strings.HasPrefix(typeStr, "[]") {
			v, err = createSliceVar(typeStr, value)
		} else {
			err = errors.New("Invalid type ")
		}
	}

	return
}

func createBaseVar(typeStr string, value interface{}, isMapKey bool) (v interface{}, err error) {
	resBytes, err := jsoniter.Marshal(value)
	if err != nil {
		return
	}

	// trim " " " if valueStr is mapKey and type in mapTypes
	resBytes = operateResBytes(typeStr, resBytes, isMapKey)

	switch typeStr {
	case "int", "int8", "int16", "int32", "int64":
		var i int64
		err = jsoniter.Unmarshal(resBytes, &i)
		v = i
	case "uint", "uint8", "uint16", "uint32", "uint64":
		var ui uint64
		err = jsoniter.Unmarshal(resBytes, &ui)
		v = ui
	case "float32", "float64":
		var f float64
		err = jsoniter.Unmarshal(resBytes, &f)
		v = f
	case "string", "types.Address":
		var s string
		err = jsoniter.Unmarshal(resBytes, &s)
		v = s
	case "bn.Number":
		var n bn.Number
		err = jsoniter.Unmarshal(resBytes, &n)
		v = n
	case "bool":
		var b bool
		err = jsoniter.Unmarshal(resBytes, &b)
		v = b
	case "byte":
		var b []byte
		resBytes = []byte(strings.Trim(string(resBytes), `"`))
		b, err = hex.DecodeString(string(resBytes[2:]))
		if err == nil {
			if len(b) == 1 {
				v = b[0]
			} else {
				err = errors.New("Type require byte, value len error ")
			}
		}
	case "types.Hash", "types.HexBytes", "types.PubKey", "[]byte":
		var bs []byte
		resBytes = []byte(strings.Trim(string(resBytes), `"`))
		bs, err = hex.DecodeString(string(resBytes[2:]))
		v = bs
	}

	return
}

func createMapVar(typeStr string, value interface{}) (v interface{}, err error) {
	keyTypeStr, valueTypeStr := getKeyValueTypeStr(typeStr)

	mapObjValue := reflect.ValueOf(value)
	ks := mapObjValue.MapKeys()

	var temp *reflect.Value
	temp, err = varFromString(typeStr)
	if err != nil {
		return
	}

	for _, k := range ks {
		v1 := mapObjValue.MapIndex(k)

		var mapKey, mapValue interface{}
		mapKey, err = createVariable(keyTypeStr, k.Interface(), true)
		if err != nil {
			return
		}

		mapValue, err = createVariable(valueTypeStr, v1.Interface(), false)
		if err != nil {
			return
		}

		temp.SetMapIndex(reflect.ValueOf(mapKey), reflect.ValueOf(mapValue))
	}
	v = temp.Interface()

	return
}

func createSliceVar(typeStr string, value interface{}) (v interface{}, err error) {
	temp := make([]interface{}, 0)
	sliceTypeStr := typeStr[2:]

	for i := 0; i < reflect.ValueOf(value).Len(); i++ {
		var v1 interface{}
		v1, err = createVariable(sliceTypeStr, reflect.ValueOf(value).Index(i).Interface(), false)
		if err != nil {
			return
		}
		temp = append(temp, v1)
	}
	v = temp

	return
}

func varFromString(typeStr string) (v *reflect.Value, err error) {
	defer funcRecover(&err)

	v = baseValue(typeStr)
	if v != nil {
		return
	}

	if strings.HasPrefix(typeStr, "[]") {
		temp := reflect.MakeSlice(typeFromString(typeStr), 0, 0)
		v = &temp
	} else if strings.HasPrefix(typeStr, "map") {
		temp := reflect.MakeMap(typeFromString(typeStr))
		v = &temp
	}

	return
}

func typeFromString(typeStr string) reflect.Type {
	typeValue := baseValue(typeStr)
	if typeValue != nil {
		return reflect.TypeOf(typeValue.Interface())
	}

	if strings.HasPrefix(typeStr, "map") {
		keyTypeStr, valueTypeStr := getKeyValueTypeStr(typeStr)

		keyType := typeFromString(keyTypeStr)
		if keyType == nil {
			return nil
		}
		valueType := typeFromString(valueTypeStr)
		if valueType == nil {
			return nil
		}

		return reflect.MapOf(keyType, valueType)
	} else if strings.HasPrefix(typeStr, "[]") {
		sliceBaseTypeStr := typeStr[2:]

		sliceBaseType := typeFromString(sliceBaseTypeStr)
		if sliceBaseType == nil {
			return nil
		}

		return reflect.SliceOf(sliceBaseType)
	}

	return nil
}

func baseValue(typeStr string) (v *reflect.Value) {
	v = new(reflect.Value)

	switch typeStr {
	case "int", "int8", "int16", "int32", "int64":
		var i int64
		*v = reflect.ValueOf(i)
	case "uint", "uint8", "uint16", "uint32", "uint64":
		var ui uint64
		*v = reflect.ValueOf(ui)
	case "float32", "float64":
		var f float64
		*v = reflect.ValueOf(f)
	case "string", "types.Address":
		var s string
		*v = reflect.ValueOf(s)
	case "bn.Number":
		var n bn.Number
		*v = reflect.ValueOf(n)
	case "bool":
		var b bool
		*v = reflect.ValueOf(b)
	case "byte":
		var b byte
		*v = reflect.ValueOf(b)
	case "types.Hash", "types.HexBytes", "types.PubKey", "[]byte":
		var sb []byte
		*v = reflect.ValueOf(sb)
	default:
		v = nil
	}

	return
}

func getKeyValueTypeStr(mapType string) (keyStr, valueStr string) {
	keyFirstIndex := strings.Index(mapType, "[")
	var prefix rune
	keyLastIndex := strings.IndexFunc(mapType, func(r rune) bool {
		if r == ']' && prefix != '[' {
			return true
		}

		prefix = r
		return false
	})

	keyStr = mapType[keyFirstIndex+1 : keyLastIndex]
	valueStr = mapType[keyLastIndex+1:]

	return
}

func operateResBytes(typeStr string, resBytes []byte, isMapKey bool) []byte {
	if isMapKey {
		var b bool
		for _, item := range mapTypes {
			if item == typeStr {
				b = true
				break
			}
		}

		if b == true {
			resBytes = []byte(strings.Trim(string(resBytes), "\""))
		}
	}

	return resBytes
}

//funcRecover recover panic by Assert
func funcRecover(errPtr *error) {
	if err := recover(); err != nil {
		msg := ""
		if errInfo, ok := err.(error); ok {
			msg = errInfo.Error()
		}

		if errInfo, ok := err.(string); ok {
			msg = errInfo
		}

		*errPtr = errors.New(msg)
	}
}
